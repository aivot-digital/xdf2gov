package gover_models

import (
	"goxdf2gov/src/utils"
	"strconv"
	"time"
)

type Field map[string]any

func NewGroupField(id string, name string, children []Field) Field {
	return Field{
		"id":         "grup_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":       ET_Group,
		"appVersion": utils.GetAppVersion(),
		"children":   children,
		"name":       name,
		"storeLink":  nil,
	}
}

func NewHeadlineField(id string, text string, small bool) Field {
	return Field{
		"id":         "hdln_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":       ET_Headline,
		"appVersion": utils.GetAppVersion(),
		"name":       text,
		"content":    text,
		"small":      small,
	}
}

func NewAlertField(id string, name string, alertType string, title string, text string) Field {
	return Field{
		"id":         "alrt_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":       ET_Alert,
		"appVersion": utils.GetAppVersion(),
		"name":       name,
		"title":      title,
		"text":       text,
		"alertType":  alertType,
	}
}

func NewTextField(id string, name string, label string, hint string, isRequired bool) Field {
	field := Field{
		"id":         "text_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":       ET_Text,
		"appVersion": utils.GetAppVersion(),
		"name":       name,
		"label":      label,
		"hint":       hint,
		"required":   isRequired,
	}
	return field
}

func NewRadioYesNo(id string, name string, label string, hint string, isRequired bool) Field {
	return Field{
		"id":         "radi_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":       ET_Radio,
		"appVersion": utils.GetAppVersion(),
		"name":       name,
		"label":      label,
		"hint":       hint,
		"required":   isRequired,
		"options": []map[string]string{
			{
				"label": "Ja",
				"value": "ja",
			},
			{
				"label": "Nein",
				"value": "nein",
			},
		},
	}
}

func NewNumberField(id string, name string, label string, hint string, isRequired bool, decimal int, suffix string) Field {
	return Field{
		"id":            "numb_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":          ET_Number,
		"appVersion":    utils.GetAppVersion(),
		"name":          name,
		"label":         label,
		"hint":          hint,
		"required":      isRequired,
		"decimalPlaces": decimal,
		"suffix":        suffix,
	}
}

func NewDateField(id string, name string, label string, hint string, isRequired bool) Field {
	return Field{
		"id":         "date_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":       ET_Date,
		"appVersion": utils.GetAppVersion(),
		"name":       name,
		"label":      label,
		"hint":       hint,
		"required":   isRequired,
	}
}

func NewFileUploadField(id string, name string, label string, hint string, isRequired bool) Field {
	return Field{
		"id":         "fupl_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":       ET_FileUpload,
		"appVersion": utils.GetAppVersion(),
		"name":       name,
		"label":      label,
		"hint":       hint,
		"required":   isRequired,
		"extensions": []string{
			"pdf",
			"png",
			"jpg",
			"jpeg",
		},
	}
}

func NewSelectField(id string, name string, label string, hint string, isRequired bool) Field {
	return Field{
		"id":         "selc_" + id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"type":       ET_Select,
		"appVersion": utils.GetAppVersion(),
		"name":       name,
		"label":      label,
		"hint":       hint,
		"required":   isRequired,
		"options":    []map[string]string{},
	}
}
