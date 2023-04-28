package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/plitn/NotesApi"
	"github.com/plitn/NotesApi/pkg/repository"
	"time"
)

const (
	salt    = "testtesttet"
	signKey = ("test")
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}
func (as *AuthService) CreateUser(user NotesApi.User) (int, error) {
	user.Password = as.genHash(user.Password)
	return as.repo.CreateUser(user)
}

type customClaim struct {
	jwt.StandardClaims
	UserId int `json:"UserId"`
}

func (as *AuthService) GenToken(username, password string) (string, error) {
	user, err := as.repo.GetUser(username, as.genHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &customClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signKey))
}

func (s *AuthService) ParseToken(token string) (int, error) {
	resultToken, err := jwt.ParseWithClaims(token, &customClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("sign in invalid")
		}
		return []byte(signKey), nil
	})
	if err != nil {
		fmt.Println("here")
		return 0, err
	}
	claims, ok := resultToken.Claims.(*customClaim)
	if !ok {
		fmt.Println("here")
		return 0, errors.New("token claims error")
	}
	return claims.UserId, nil
}

func (s *AuthService) genHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
