package processors

import (
	"fmt"
	"goxdf2gov/src/models/gover_models"
	"goxdf2gov/src/models/xdf2_models"
	"goxdf2gov/src/utils"
	"regexp"
	"strconv"
	"time"
)

var anzahlRegEx = regexp.MustCompile("^(0|([1-9][0-9]*)):(\\*|(0|([1-9][0-9]*)))$")

func processStruktur(struktur xdf2_models.Struktur, depth int) ([]gover_models.Field, error) {
	var fields []gover_models.Field

	anzahlParts := anzahlRegEx.FindStringSubmatch(struktur.Anzahl)
	minAnzahlStr := anzahlParts[1]
	maxAnzahlStr := anzahlParts[3]

	minAnzahl, err := strconv.Atoi(minAnzahlStr)
	if err != nil {
		return nil, err
	}

	maxAnzahl := 1
	if maxAnzahlStr == "*" {
		maxAnzahl = -1
	} else {
		maxAnzahl, err = strconv.Atoi(maxAnzahlStr)
		if err != nil {
			return nil, err
		}
	}

	isRequired := minAnzahl > 0
	isReplicating := maxAnzahl == -1 || maxAnzahl > 1

	if struktur.Enthaelt.Datenfeldgruppe.Identifikation.Id != "" {
		_fields, err := processDatenfeldgruppe(struktur.Enthaelt.Datenfeldgruppe, isRequired, depth)
		if err != nil {
			return nil, err
		}
		fields = append(fields, _fields...)
	}

	if struktur.Enthaelt.Datenfeld.Identifikation.Id != "" {
		_fields, err := processDatenfeld(struktur.Enthaelt.Datenfeld, isRequired)
		if err != nil {
			return nil, err
		}
		fields = append(fields, _fields...)
	}

	isUploadContainer := len(fields) == 1 && fields[0]["type"] == gover_models.ET_FileUpload

	if isUploadContainer {
		fields[0]["isMultifile"] = minAnzahl > 0 || maxAnzahl > 1 || maxAnzahl == -1
		if minAnzahl > 0 {
			fields[0]["minFiles"] = minAnzahl
			fields[0]["required"] = true
		}
		if maxAnzahl > 1 {
			fields[0]["maxFiles"] = maxAnzahl
			fields[0]["required"] = true
		}
	} else if isReplicating && len(fields) > 0 {
		firstChild := fields[0]
		label := firstChild["label"]
		if label == nil || label == "" {
			label = firstChild["content"]
		}
		if label == nil || label == "" {
			label = firstChild["title"]
		}
		if label == nil || label == "" {
			label = "Eintrag"
		}

		replicatingContainer := gover_models.Step{
			"id":               "repc_" + strconv.FormatInt(time.Now().UnixNano(), 10),
			"type":             gover_models.ET_ReplicatingContainer,
			"appVersion":       utils.GetAppVersion(),
			"children":         fields,
			"required":         isRequired || minAnzahl > 0,
			"label":            label,
			"hint":             firstChild["hint"],
			"headlineTemplate": fmt.Sprintf("%s #", label),
			"addLabel":         fmt.Sprintf("%s hinzufügen", label),
			"removeLabel":      fmt.Sprintf("%s löschen", label),
		}
		if minAnzahl > 0 {
			replicatingContainer["minimumRequiredSets"] = minAnzahl
		}
		if maxAnzahl > 1 {
			replicatingContainer["maximumSets"] = maxAnzahl
		}
		return []gover_models.Field{
			replicatingContainer,
		}, nil
	}

	return fields, nil
}
