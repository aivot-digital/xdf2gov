package gover_models

type Todo struct {
	StepId  string `json:"stepId"`
	FieldId string `json:"fieldId"`
	Text    string `json:"text"`
	Done    bool   `json:"done"`
}
