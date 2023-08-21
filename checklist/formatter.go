package checklist

type ChecklistFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatChecklist(checklist Checklist) ChecklistFormatter {
	formatter := ChecklistFormatter{
		ID:   checklist.ID,
		Name: checklist.Name,
	}
	return formatter
}

func FormatChecklists(Checklists []Checklist) []ChecklistFormatter {

	ChecklistFormatter := []ChecklistFormatter{}

	for _, Checklist := range Checklists {
		checklistFormatter := FormatChecklist(Checklist)
		ChecklistFormatter = append(ChecklistFormatter, checklistFormatter)

	}

	return ChecklistFormatter
}
