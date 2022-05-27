package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Ckala62rus/rest-go"
	"github.com/Ckala62rus/rest-go/pkg/repository"
)

const salt = "qweasdzxc" // хорошая практика добавлять соль к паролю

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user rest.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
