package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Ckala62rus/rest-go"
	"github.com/Ckala62rus/rest-go/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "qweasdzxc" // хорошая практика добавлять соль к паролю
	signingKey = "qweasadasfasfasfsfafasf"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))

	if user.Id == 0 {
		return "User not found", err
	}

	// logrus.Fatal(user.Id)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
