package dto

import (
	"github.com/AlexeevNikita/joint-sphere/internal/entities"
)

type GetUserResponse struct {
	ID             int64  `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Country        string `json:"country"`
	Birthdate      string `json:"birthdate"`
	AdditionalInfo string `json:"additionalInfo"`
}

func NewGetUserResponse(user *entities.User) *GetUserResponse {
	response := &GetUserResponse{
		ID:             user.ID,
		Email:          user.Email,
		Username:       user.Username,
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     user.Patronymic,
		Country:        user.Country,
		Birthdate:      user.Birthdate.String(),
		AdditionalInfo: user.AdditionalInfo,
	}

	return response
}
