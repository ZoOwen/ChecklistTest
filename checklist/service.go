package checklist

type Service interface {
	GetChecklist() ([]Checklist, error)
	CreateChecklist(input CreateChecklistInput) (Checklist, error)
	DeleteChecklist(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetChecklist() ([]Checklist, error) {

	checklists, err := s.repository.FindAll()
	if err != nil {
		return checklists, err
	}

	return checklists, nil
}

func (s *service) CreateChecklist(input CreateChecklistInput) (Checklist, error) {
	Checklist := Checklist{}
	Checklist.Name = input.Name

	newChecklist, err := s.repository.Save(Checklist)
	if err != nil {
		return newChecklist, err
	}

	return newChecklist, nil
}

func (s *service) DeleteChecklist(ID int) error {
	err := s.repository.DeleteChecklist(ID)
	if err != nil {
		return err
	}

	return nil
}
