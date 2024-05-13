package service

import (
	"github.com/Insid1/go-auth-user/internal/entity"
)

type User interface {
	Create(*entity.User) (string, error)
	Get(string) (*entity.User, error)
	Update(*entity.User) (*entity.User, error)
	Delete(string) error
}
