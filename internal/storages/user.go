package storages

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AlexeevNikita/joint-sphere/internal/entities"
)

type userStorage struct {
	db *sql.DB
}

func NewUserStorage(
	db *sql.DB,
) (entities.UserStorage, error) {

	s := &userStorage{db: db}

	return s, nil
}

func (s *userStorage) Store(ctx context.Context, user *entities.User) error {
	const query = "INSERT INTO public.user (email, username, surname, name, patronymic, country, birthdate, additional_info) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := s.db.ExecContext(ctx, query, user.Email, user.Username, user.Surname, user.Name, user.Patronymic, user.Country, user.Birthdate, user.AdditionalInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("problem while trying to create user: %s", err.Error()))
	}

	return nil
}

func (s *userStorage) Get(ctx context.Context, ID int64) (*entities.User, error) {
	const query = "SELECT id, email, username, surname, name, patronymic, country, birthdate, additional_info FROM public.user WHERE id=$1"
	rows, err := s.db.QueryContext(ctx, query, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("problem while trying to get user by id: %s", err.Error()))
	}

	defer rows.Close()

	user := &entities.User{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Surname, &user.Name, &user.Patronymic, &user.Country, &user.Birthdate, &user.AdditionalInfo); err != nil {
			return nil, errors.New(fmt.Sprintf("problem while trying to retrieve user by id: %s", err.Error()))
		}
	}

	return user, nil
}

func (s *userStorage) Patch(ctx context.Context, userData *entities.User) error {
	const query = "UPDATE public.user SET email=$1, username=$2, surname=$3, name=$4, patronymic=$5, country=$6, birthdate=$7, additional_info=$8 WHERE id=$9"
	_, err := s.db.ExecContext(ctx, query, userData.Email, userData.Username, userData.Surname, userData.Name, userData.Patronymic, userData.Country, userData.Birthdate, userData.AdditionalInfo, userData.ID)
	if err != nil {
		return errors.New(fmt.Sprintf("problem while trying to update user: %s", err.Error()))
	}

	return nil
}

func (s *userStorage) Delete(ctx context.Context, ID int64) error {
	const query = "DELETE FROM public.user WHERE id=$1"
	_, err := s.db.ExecContext(ctx, query, ID)
	if err != nil {
		return errors.New(fmt.Sprintf("problem while trying to delete user by id: %s", err.Error()))
	}

	return nil
}
