package checklist

type GetChecklistDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateChecklistInput struct {
	Name string `json:"name" binding:"required"`
}
