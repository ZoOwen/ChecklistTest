package checklistItem

type ChecklistItemFormatter struct {
	ID          int    `json:"id"`
	ChecklistId int    `json:"checklist_id"`
	ItemName    string `json:"item_name"`
}

func FormatChecklist(checklistItem ChecklistItem) ChecklistItemFormatter {
	formatter := ChecklistItemFormatter{
		ID:          checklistItem.ID,
		ChecklistId: checklistItem.ChecklistId,
		ItemName:    checklistItem.ItemName,
	}
	return formatter
}

func FormatChecklists(Checklists []ChecklistItem) []ChecklistItemFormatter {

	ChecklistsFormatter := []ChecklistItemFormatter{}

	for _, Checklist := range Checklists {
		ChecklistFormatter := FormatChecklist(Checklist)
		ChecklistsFormatter = append(ChecklistsFormatter, ChecklistFormatter)

	}

	return ChecklistsFormatter
}
