package service

import (
	"task-manager/internal/entity"
	"task-manager/internal/repository"
)

type UserService interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type userService struct {
	repo repository.UserRepoitory
}

func NewUserService(r repository.UserRepoitory) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(user *entity.User) (*entity.User, error) {
	return s.repo.Create(user)
}

func (s *userService) GetUserByEmail(email string) (*entity.User, error) {
	return s.repo.FindByEmail(email)
}
