package checklistItem

import "fmt"

type Service interface {
	GetChecklistItem(checklistID int) ([]ChecklistItem, error)
	CreateChecklistItem(input CreateChecklistInput, checklistIDInt int) (ChecklistItem, error)
	DeleteChecklist(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetChecklistItem(checklistID int) ([]ChecklistItem, error) {

	checklists, err := s.repository.FindAll(checklistID)
	if err != nil {
		return checklists, err
	}
	fmt.Print("ini data checklists", checklists)
	return checklists, nil
}

func (s *service) CreateChecklistItem(input CreateChecklistInput, checklistIDInt int) (ChecklistItem, error) {
	checklist := ChecklistItem{}
	checklist.ItemName = input.ItemName
	checklist.ChecklistId = checklistIDInt
	newchecklist, err := s.repository.Save(checklist)
	if err != nil {
		return newchecklist, err
	}

	return newchecklist, nil
}

func (s *service) DeleteChecklist(ID int) error {
	err := s.repository.DeleteChecklist(ID)
	if err != nil {
		return err
	}

	return nil
}
