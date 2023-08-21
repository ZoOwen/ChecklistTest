package checklistItem

import "fmt"

type Service interface {
	GetChecklistItem(checklistID int) ([]ChecklistItem, error)
	CreateChecklistItem(input CreateChecklistInput, checklistIDInt int) (ChecklistItem, error)
	GetChecklistsItemId(ID int) (ChecklistItem, error)
	UpdateChecklistItem(ID int, itemName string) error
	DeleteChecklistItem(ID int) error
	UpdateStatusChecklistItem(ID int) error
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

func (s *service) GetChecklistsItemId(ID int) (ChecklistItem, error) {
	data, err := s.repository.GetChecklistsItemId(ID)
	if err != nil {
		return ChecklistItem{}, err
	}

	return data, nil
}

func (s *service) UpdateChecklistItem(ID int, itemName string) error {
	checklistItem, err := s.repository.GetChecklistsItemId(ID)
	if err != nil {
		return err
	}

	checklistItem.ItemName = itemName

	err = s.repository.UpdateChecklistItem(checklistItem)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteChecklistItem(ID int) error {
	err := s.repository.DeleteChecklistItem(ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateStatusChecklistItem(ID int) error {
	checklistItem, err := s.repository.GetChecklistsItemId(ID)
	if err != nil {
		return err
	}

	checklistItem.Status = "done"

	err = s.repository.UpdateChecklistItem(checklistItem)
	if err != nil {
		return err
	}

	return nil
}
