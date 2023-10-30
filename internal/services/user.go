package services

import (
	"context"
	"errors"
	"github.com/AlexeevNikita/joint-sphere/internal/entities"
)

type userService struct {
	userStorage entities.UserStorage
}

func NewUserService(
	userStorage entities.UserStorage,
) (entities.UserService, error) {
	if userStorage == nil {
		return nil, errors.New("userStorage is nil")
	}

	srv := &userService{userStorage: userStorage}

	return srv, nil
}

func (srv *userService) Create(ctx context.Context, user *entities.User) error {
	if err := srv.userStorage.Store(ctx, user); err != nil {
		return err
	}

	return nil
}

func (srv *userService) Get(ctx context.Context, ID int64) (*entities.User, error) {
	user, err := srv.userStorage.Get(ctx, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *userService) Update(ctx context.Context, userData *entities.User) error {
	if err := srv.userStorage.Patch(ctx, userData); err != nil {
		return err
	}

	return nil
}

func (srv *userService) Delete(ctx context.Context, ID int64) error {
	if err := srv.userStorage.Delete(ctx, ID); err != nil {
		return err
	}

	return nil
}
