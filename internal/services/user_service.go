package services

import (
	"auth-template/internal/models"
	"errors"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetById(id int) (*models.User, error)
	Create(name, email, password string) (*models.User, error)
}

type userService struct {
	users  []models.User
	nextID int
}

func NewUserService() UserService {
	return &userService{
		nextID: 1,
	}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.users, nil
}

func (s *userService) GetById(id int) (*models.User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *userService) Create(name, email, password string) (*models.User, error) {
	user := models.User{
		ID:       s.nextID,
		Name:     name,
		Email:    email,
		Password: password,
	}
	s.nextID++
	s.users = append(s.users, user)
	return &user, nil
}
