package checklistItem

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(checklistID int) ([]ChecklistItem, error)
	Save(checklist ChecklistItem) (ChecklistItem, error)
	GetChecklistsItemId(ID int) (ChecklistItem, error)
	UpdateChecklistItem(checklistItem ChecklistItem) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(checklistID int) ([]ChecklistItem, error) {
	var checklistItems []ChecklistItem
	err := r.db.Where("checklist_id = ?", checklistID).Find(&checklistItems).Error
	if err != nil {
		return checklistItems, err
	}

	fmt.Println("ini id", checklistID, "ini ceklis item", checklistItems)
	return checklistItems, nil
}

func (r *repository) Save(checklist ChecklistItem) (ChecklistItem, error) {
	err := r.db.Create(&checklist).Error
	if err != nil {
		return checklist, err
	}
	return checklist, nil
}

func (r *repository) GetChecklistsItemId(ID int) (ChecklistItem, error) {
	checklistItem := ChecklistItem{}

	err := r.db.First(&checklistItem, ID).Error
	if err != nil {
		return checklistItem, err
	}

	return checklistItem, nil
}

func (r *repository) UpdateChecklistItem(checklistItem ChecklistItem) error {
	err := r.db.Save(&checklistItem).Error
	if err != nil {
		return err
	}

	return nil
}
