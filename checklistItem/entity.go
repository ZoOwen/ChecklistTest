package checklistItem

type ChecklistItem struct {
	ID          int    `json:"id"`
	ChecklistId int    `json:"checklist_id"`
	ItemName    string `json:"item_name"`
	Status      string `json:"status"`
}
