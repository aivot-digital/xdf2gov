package processors

import (
	"goxdf2gov/src/models/gover_models"
	"goxdf2gov/src/models/xdf2_models"
	"goxdf2gov/src/utils"
	"strings"
)

func processDatenfeldgruppe(datenfeldgruppe xdf2_models.Datenfeldgruppe, isRequired bool, depth int) ([]gover_models.Field, error) {
	children := make([]gover_models.Field, 0)

	if depth >= 1 {
		children = append(
			children,
			gover_models.NewHeadlineField(
				datenfeldgruppe.Identifikation.Id,
				utils.CleanString(datenfeldgruppe.BezeichnungEingabe),
				depth >= 2,
			),
		)
	}

	for _, struktur := range datenfeldgruppe.Struktur {
		f, err := processStruktur(struktur, depth+1)
		if err != nil {
			return nil, err
		}
		children = append(children, f...)
	}

	id := strings.ToLower(datenfeldgruppe.Identifikation.Id)
	name := utils.CleanString(datenfeldgruppe.Name)

	return []gover_models.Field{
		gover_models.NewGroupField(id, name, children),
	}, nil
}
