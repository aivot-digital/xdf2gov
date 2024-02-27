package processors

import (
	"goxdf2gov/src/models/gover_models"
	"goxdf2gov/src/models/xdf2_models"
	"goxdf2gov/src/utils"
	"strconv"
	"strings"
	"time"
)

func ProcessStammdatenschema(stammdatenschema xdf2_models.Stammdatenschema) (gover_models.Form, error) {
	id := strings.ToLower(stammdatenschema.Identifikation.Id)
	name := utils.CleanString(stammdatenschema.Name)

	var description string
	if stammdatenschema.Beschreibung != "" {
		description = stammdatenschema.Beschreibung
	} else if stammdatenschema.Definition != "" {
		description = stammdatenschema.Definition
	} else if stammdatenschema.BezeichnungEingabe != "" {
		description = stammdatenschema.BezeichnungEingabe
	} else {
		description = "Unbenannter Abschnitt"
	}
	description = utils.CleanString(description)

	var steps []gover_models.Step
	for _, stepStruktur := range stammdatenschema.Struktur {
		fields, err := processStruktur(stepStruktur, 0)
		if err != nil {
			return nil, err
		}

		var newStep gover_models.Step
		if stepStruktur.Enthaelt.Datenfeldgruppe.Identifikation.Id != "" {
			newStep = gover_models.Step{
				"id":         "step_" + strings.ToLower(stepStruktur.Enthaelt.Datenfeldgruppe.Identifikation.Id) + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
				"type":       gover_models.ET_Step,
				"title":      utils.CleanString(stepStruktur.Enthaelt.Datenfeldgruppe.Name),
				"appVersion": utils.GetAppVersion(),
				"children":   fields,
				"icon":       "arrow",
			}
		} else if stepStruktur.Enthaelt.Datenfeld.Identifikation.Id != "" {
			newStep = gover_models.Step{
				"id":         "step_" + strings.ToLower(stepStruktur.Enthaelt.Datenfeld.Identifikation.Id) + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
				"type":       gover_models.ET_Step,
				"title":      utils.CleanString(stepStruktur.Enthaelt.Datenfeld.Name),
				"appVersion": utils.GetAppVersion(),
				"children":   fields,
				"icon":       "arrow",
			}
		}

		steps = append(steps, newStep)
	}

	return gover_models.Form{
		"id":                         0,
		"slug":                       "",
		"version":                    "1.0.0",
		"headline":                   name,
		"theme":                      nil,
		"title":                      name,
		"status":                     0,
		"developingDepartment":       nil,
		"managingDepartment":         nil,
		"responsibleDepartment":      nil,
		"created":                    time.Now().Format("2006-01-02T15:04:05.000Z"),
		"updated":                    time.Now().Format("2006-01-02T15:04:05.000Z"),
		"openSubmissions":            0,
		"inProgressSubmissions":      0,
		"totalSubmissions":           0,
		"destination":                nil,
		"legalSupportDepartment":     nil,
		"technicalSupportDepartment": nil,
		"imprintDepartment":          nil,
		"privacyDepartment":          nil,
		"accessibilityDepartment":    nil,
		"customerAccessHours":        0,
		"submissionDeletionWeeks":    0,
		"root": map[string]any{
			"type":            0,
			"id":              "root_" + id,
			"appVersion":      utils.GetAppVersion(),
			"name":            "",
			"testProtocolSet": nil,
			"isVisible":       nil,
			"patchElement":    nil,
			"headline":        name,
			"tabTitle":        "",
			"children":        steps,
			"expiring":        nil,
			"accessLevel":     nil,
			"privacyText":     "Bitte beachten Sie die {privacy}Hinweise zum Datenschutz{/privacy}.",
			"introductionStep": map[string]any{
				"type":                17,
				"id":                  "intr_" + id,
				"appVersion":          utils.GetAppVersion(),
				"name":                "",
				"testProtocolSet":     nil,
				"isVisible":           nil,
				"patchElement":        nil,
				"initiativeName":      nil,
				"initiativeLogoLink":  nil,
				"initiativeLink":      nil,
				"teaserText":          description,
				"organization":        nil,
				"eligiblePersons":     nil,
				"supportingDocuments": nil,
				"documentsToAttach":   nil,
				"expectedCosts":       nil,
			},
			"summaryStep": map[string]any{
				"type":            19,
				"id":              "summ_" + id,
				"appVersion":      utils.GetAppVersion(),
				"name":            "",
				"testProtocolSet": nil,
				"isVisible":       nil,
				"patchElement":    nil,
			},
			"submitStep": map[string]any{
				"type":               18,
				"id":                 "subm_" + id,
				"appVersion":         utils.GetAppVersion(),
				"name":               "",
				"testProtocolSet":    nil,
				"isVisible":          nil,
				"patchElement":       nil,
				"textPreSubmit":      "Sie können Ihren Antrag nun verbindlich bei der zuständigen/bewirtschaftenden Stelle einreichen. Nach der Einreichung können Sie sich den Antrag für Ihre Unterlagen herunterladen oder zusenden lassen.",
				"textPostSubmit":     "Sie können Ihren Antrag herunterladen oder sich per E-Mail zuschicken lassen. Wir empfehlen Ihnen, den Antrag anschließend zu Ihren Unterlagen zu nehmen.",
				"textProcessingTime": nil,
				"documentsToReceive": nil,
			},
		},
	}, nil
}
