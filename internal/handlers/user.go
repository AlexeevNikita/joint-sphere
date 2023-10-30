package handlers

import (
	"context"
	"errors"
	"github.com/AlexeevNikita/joint-sphere/internal/dto"
	"github.com/AlexeevNikita/joint-sphere/internal/entities"
	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strconv"
	"time"
)

type userHandler struct {
	userService entities.UserService
}

func NewUserHandler(
	userService entities.UserService,
) (entities.UserHandler, error) {

	if userService == nil {
		return nil, errors.New("userService is nil")
	}

	h := &userHandler{userService: userService}

	return h, nil
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	birthdateString := r.Form.Get("birthdate")
	birthdate, err := time.Parse("2006-01-02", birthdateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &entities.User{
		Email:          r.Form.Get("email"),
		Username:       r.Form.Get("username"),
		Surname:        r.Form.Get("surname"),
		Name:           r.Form.Get("name"),
		Patronymic:     r.Form.Get("patronymic"),
		Country:        r.Form.Get("country"),
		Birthdate:      birthdate,
		AdditionalInfo: r.Form.Get("additional_info"),
	}

	err = h.userService.Create(context.Background(), user)
	switch err {
	case nil:
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *userHandler) Read(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.Get(context.Background(), ID)
	switch err {
	case nil:
		w.WriteHeader(200)
		if err := jsoniter.NewEncoder(w).Encode(dto.NewGetUserResponse(user)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	birthdateString := r.Form.Get("birthdate")
	birthdate, err := time.Parse("2006-01-02", birthdateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &entities.User{
		ID:             ID,
		Email:          r.Form.Get("email"),
		Username:       r.Form.Get("username"),
		Surname:        r.Form.Get("surname"),
		Name:           r.Form.Get("name"),
		Patronymic:     r.Form.Get("patronymic"),
		Country:        r.Form.Get("country"),
		Birthdate:      birthdate,
		AdditionalInfo: r.Form.Get("additional_info"),
	}

	err = h.userService.Update(context.Background(), user)
	switch err {
	case nil:
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.Delete(context.Background(), ID)
	switch err {
	case nil:
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
