package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Username = input.Username
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Email = input.Email
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	Username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(Username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No User Found On That Email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *service) GetUserByID(ID int) (User, error) {

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No User Found With That ID")
	}

	return user, nil

}
