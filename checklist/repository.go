package checklist

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Checklist, error)
	Save(checklist Checklist) (Checklist, error)
	DeleteChecklist(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Checklist, error) {
	var checklists []Checklist
	err := r.db.Find(&checklists).Error
	if err != nil {
		return checklists, err
	}
	return checklists, nil
}

func (r *repository) Save(checklist Checklist) (Checklist, error) {
	err := r.db.Create(&checklist).Error
	if err != nil {
		return checklist, err
	}
	return checklist, nil
}

func (r *repository) DeleteChecklist(ID int) error {
	checklist := Checklist{}

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
