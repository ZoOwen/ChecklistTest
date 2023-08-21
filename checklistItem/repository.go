package checklistItem

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(checklistID int) ([]ChecklistItem, error)
	Save(checklist ChecklistItem) (ChecklistItem, error)
	DeleteChecklist(ID int) error
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

func (r *repository) DeleteChecklist(ID int) error {
	checklist := ChecklistItem{}

	err := r.db.First(&checklist, ID).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&checklist).Error
	if err != nil {
		return err
	}

	return nil
}
