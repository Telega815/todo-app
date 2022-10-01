package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/telega815/todo-app"
	"github.com/telega815/todo-app/pkg/repository"
)

const salt = "f7H34T6t34AtfNu"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error)  {
	user.Password = generatePassHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePassHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
