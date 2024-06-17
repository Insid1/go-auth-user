package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Insid1/go-auth-user/auth-service/internal/model"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository"
	"github.com/Insid1/go-auth-user/auth-service/pkg/user_v1"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Ctx    context.Context
	JWTKey string

	UserRepository repository.User
}

func (s *Service) Login(data *model.Login) (string, error) {
	var usr *user_v1.User

	usr, err := s.UserRepository.Get(data.ID, data.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.GetPassHash()), []byte(data.Password))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error: Password is invalid. %s", err))
	}

	token, err := s.generateToken(usr.GetId(), usr.GetEmail(), usr.GetPassHash())
	if err != nil {
		return "", err
	}

	return token, nil
}
func (s *Service) Register(*model.Register) (string, error) {
	return "", nil
}
func (s *Service) Logout(string) (bool, error) {
	return false, nil
}

func (s *Service) generateToken(id string, email string, passHash string) (string, error) {
	// Генерируем полезные данные, которые будут храниться в токене
	payload := jwt.MapClaims{
		"sub":   id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(s.JWTKey + passHash))
}