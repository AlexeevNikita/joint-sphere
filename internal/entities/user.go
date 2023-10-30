package entities

import (
	"context"
	"net/http"
	"time"
)

type User struct {
	ID             int64
	Email          string
	Username       string
	Surname        string
	Name           string
	Patronymic     string
	Country        string
	Birthdate      time.Time
	AdditionalInfo string
}

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type UserService interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, ID int64) (*User, error)
	Update(ctx context.Context, userData *User) error
	Delete(ctx context.Context, ID int64) error
}

type UserStorage interface {
	Store(ctx context.Context, user *User) error
	Get(ctx context.Context, ID int64) (*User, error)
	Patch(ctx context.Context, userData *User) error
	Delete(ctx context.Context, ID int64) error
}
