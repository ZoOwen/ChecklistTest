package checklistItem

type GetChecklistDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateChecklistInput struct {
	ItemName string `json:"item_name" binding:"required"`
}
