package processors

import (
	"encoding/json"
	"fmt"
	"goxdf2gov/src/models"
	"goxdf2gov/src/models/gover_models"
	"goxdf2gov/src/models/xdf2_models"
	"goxdf2gov/src/utils"
	"strconv"
	"strings"
)

func processDatenfeld(datenfeld xdf2_models.Datenfeld, isRequired bool) ([]gover_models.Field, error) {
	var fields []gover_models.Field

	_id := utils.CleanString(datenfeld.Identifikation.Id)
	id := strings.ToLower(_id)
	name := utils.CleanString(datenfeld.Name)

	switch datenfeld.Feldart.Code {
	case "input":
		label := utils.CleanString(datenfeld.BezeichnungEingabe)
		hint := utils.CleanString(datenfeld.HilfetextEingabe)

		praezisierung, err := parsePraezisierung(datenfeld.Praezisierung)
		if err != nil {
			fields = append(
				fields,
				gover_models.NewAlertField(
					id,
					"--- Fehler ---",
					"error",
					"Fehler beim Laden der Präzisierung",
					fmt.Sprintf("Die Präzisierung des Feldes %s (%s) konnte nicht geladen werden: Präzisiserung: \"%s\" - Fehler: %s", name, _id, datenfeld.Praezisierung, err.Error()),
				),
			)
			praezisierung = nil
		}

		switch datenfeld.Datentyp.Code {
		case "text":
			field := gover_models.NewTextField(id, name, label, hint, isRequired)

			if praezisierung != nil {
				if praezisierung.MinLength != nil {
					field["minCharacters"] = *praezisierung.MinLength
					if *praezisierung.MinLength > 255 {
						field["isMultiline"] = true
					}
				}

				if praezisierung.MaxLength != nil {
					field["maxCharacters"] = *praezisierung.MaxLength
					if *praezisierung.MaxLength > 255 {
						field["isMultiline"] = true
					}
				}

				if praezisierung.Pattern != nil {
					field["validate"] = map[string]any{
						"requirements": "",
						"conditionSet": map[string]any{
							"operator": 0,
							"conditions": []map[string]any{
								{
									"reference":             field["id"],
									"operator":              12,
									"value":                 *praezisierung.Pattern,
									"conditionUnmetMessage": fmt.Sprintf("Bitte geben Sie einen Wert an der dem Muster \"%s\" entspricht.", *praezisierung.Pattern),
								},
							},
							"conditionsSets": []map[string]any{},
						},
					}
				}
			}

			fields = append(fields, field)
			break
		case "bool":
			fields = append(fields, gover_models.NewRadioYesNo(id, name, label, hint, isRequired))
			break
		case "num":
			field := gover_models.NewNumberField(id, name, label, hint, isRequired, 2, "")
			validate := getValidateFunc(field["id"], praezisierung)
			if validate != nil {
				field["validate"] = validate
			}
			fields = append(fields, field)
			break
		case "num_int":
			field := gover_models.NewNumberField(id, name, label, hint, isRequired, 0, "")
			validate := getValidateFunc(field["id"], praezisierung)
			if validate != nil {
				field["validate"] = validate
			}
			fields = append(fields, field)
			break
		case "num_currency":
			field := gover_models.NewNumberField(id, name, label, hint, isRequired, 2, "€")
			validate := getValidateFunc(field["id"], praezisierung)
			if validate != nil {
				field["validate"] = validate
			}
			fields = append(fields, field)
			break
		case "file":
			fields = append(fields, gover_models.NewFileUploadField(id, name, label, hint, isRequired))
			break
		case "date":
			fields = append(fields, gover_models.NewDateField(id, name, label, hint, isRequired))
			break
		default:
			fields = append(
				fields,
				gover_models.NewAlertField(
					id,
					"--- Fehler ---",
					"error",
					"Nicht unterstützter Datentyp",
					fmt.Sprintf("Datentype %s für die Feldart %s für das Feld mit der Id %s wird nicht unterstützt", datenfeld.Datentyp.Code, datenfeld.Feldart.Code, _id),
				),
			)
		}
		break
	case "select":
		label := utils.CleanString(datenfeld.BezeichnungEingabe)
		hint := utils.CleanString(datenfeld.HilfetextEingabe)
		field := gover_models.NewSelectField(id, name, label, hint, isRequired)

		codeListUrn := utils.CleanString(datenfeld.CodelisteReferenz.GenericodeIdentification.CanonicalVersionUri)
		codelist, err := utils.FetchCodeList(codeListUrn)
		if err != nil {
			fields = append(
				fields,
				gover_models.NewAlertField(
					id,
					"--- Fehler ---",
					"error",
					"CodeListe konnte nicht geladen werden",
					fmt.Sprintf("Die Code-Liste mit der URN %s für das Feld mit der Id %s konnte nicht geladen werden.", codeListUrn, _id),
				),
			)
		} else {
			field["options"] = codelist
		}

		fields = append(fields, field)
		break
	case "label":
		fields = append(
			fields,
			gover_models.NewAlertField(
				id,
				utils.CleanString(datenfeld.Name),
				"info",
				utils.CleanString(datenfeld.BezeichnungEingabe),
				utils.CleanString(datenfeld.Inhalt),
			),
		)
		break
	default:
		fields = append(
			fields,
			gover_models.NewAlertField(
				id,
				"--- Fehler ---",
				"error",
				"Nicht unterstützte Feldart",
				fmt.Sprintf("Feldart %s für das Feld mit der Id %s wird nicht unterstützt", datenfeld.Feldart.Code, _id),
			),
		)
	}

	return fields, nil
}

func parsePraezisierung(raw string) (*models.Praezisierung, error) {
	if raw == "" {
		return nil, nil
	}

	praezisiserung := models.Praezisierung{}

	var praezisierungRaw map[string]any
	err := json.Unmarshal([]byte(raw), &praezisierungRaw)
	if err != nil {
		return nil, err
	}

	allowedKeys := []string{
		"minLength",
		"maxLength",
		"pattern",
		"minValue",
		"maxValue",
	}

	isKeyAllowed := func(key string) bool {
		for _, allowedKey := range allowedKeys {
			if key == allowedKey {
				return true
			}
		}
		return false
	}

	for key, _ := range praezisierungRaw {
		if !isKeyAllowed(key) {
			return nil, fmt.Errorf("Der Schlüssel %s in der Präsiziserung ist nicht erlaubt.", key)
		}
	}

	praezisiserung.MinLength, err = parseIntValue(praezisierungRaw, "minLength")
	if err != nil {
		return nil, err
	}
	praezisiserung.MaxLength, err = parseIntValue(praezisierungRaw, "maxLength")
	if err != nil {
		return nil, err
	}
	praezisiserung.MinValue, err = parseIntValue(praezisierungRaw, "minValue")
	if err != nil {
		return nil, err
	}
	praezisiserung.MaxValue, err = parseIntValue(praezisierungRaw, "maxValue")
	if err != nil {
		return nil, err
	}
	praezisiserung.Pattern, err = parseStringValue(praezisierungRaw, "pattern")
	if err != nil {
		return nil, err
	}

	return &praezisiserung, nil
}

func parseIntValue(mapRaw map[string]any, source string) (*int, error) {
	valueRaw := mapRaw[source]

	if valueRaw == nil {
		return nil, nil
	}

	value := -1
	var ok bool
	var err error

	if value, ok = valueRaw.(int); !ok {
		var valueStr string
		if valueStr, ok = valueRaw.(string); ok {
			value, err = strconv.Atoi(valueStr)
			if err != nil {
				return nil, fmt.Errorf("Der Wert \"%s\" für die Präsiziserung %s ist keine gültige Zahl.", valueStr, source)
			}
		} else {
			return nil, fmt.Errorf("Der Wert \"%v\" für die Präsiziserung %s ist keine gültige Zahl.", valueRaw, source)
		}
	}

	return &value, nil
}

func parseStringValue(mapRaw map[string]any, source string) (*string, error) {
	valueRaw := mapRaw[source]

	if valueRaw == nil {
		return nil, nil
	}

	if value, ok := valueRaw.(string); !ok {
		return nil, fmt.Errorf("Der Wert \"%v\" für die Präsiziserung %s ist kein gültiger Text.", valueRaw, source)
	} else {
		return &value, nil
	}
}

func getValidateFunc(fieldId any, praezisierung *models.Praezisierung) map[string]any {
	if praezisierung != nil && (praezisierung.MinValue != nil || praezisierung.MaxValue != nil) {
		var conditions []map[string]any

		if praezisierung.MinValue != nil {
			conditions = append(conditions, map[string]any{
				"reference":             fieldId,
				"operator":              5,
				"value":                 fmt.Sprintf("%d", *praezisierung.MinValue),
				"conditionUnmetMessage": fmt.Sprintf("Bitte geben Sie einen Wert größer oder gleich %d ein.", *praezisierung.MinValue),
			})
		}

		if praezisierung.MaxValue != nil {
			conditions = append(conditions, map[string]any{
				"reference":             fieldId,
				"operator":              3,
				"value":                 fmt.Sprintf("%d", *praezisierung.MaxValue),
				"conditionUnmetMessage": fmt.Sprintf("Bitte geben Sie einen Wert kleiner oder gleich %d ein.", *praezisierung.MaxValue),
			})
		}

		return map[string]any{
			"requirements": "",
			"conditionSet": map[string]any{
				"operator":       0,
				"conditions":     conditions,
				"conditionsSets": []any{},
			},
		}
	}
	return nil
}
