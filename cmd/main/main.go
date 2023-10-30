package main

import (
	"errors"
	"fmt"
	"github.com/AlexeevNikita/joint-sphere/internal/entities"
	"github.com/AlexeevNikita/joint-sphere/internal/handlers"
	"github.com/AlexeevNikita/joint-sphere/internal/resthttp"
	"github.com/AlexeevNikita/joint-sphere/internal/services"
	"github.com/AlexeevNikita/joint-sphere/internal/storages"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var (
	sqlxConnectDataTemplate = "user=%s password=%s dbname=postgres sslmode=disable"
)

func main() {
	cfg := &entities.Config{
		Database: struct {
			User     string
			Password string
		}{User: os.Getenv("database_user"), Password: os.Getenv("database_password")}}

	if err := run(cfg); err != nil {
		log.Fatalln("problem while trying to start application", err)
	}
}

func run(cfg *entities.Config) error {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(sqlxConnectDataTemplate, cfg.Database.User, cfg.Database.Password))
	if err != nil {
		log.Fatalln("error occurred while trying to connect to postgres", err)
	}

	userStorage, err := storages.NewUserStorage(db.DB)
	if err != nil {
		return errors.New(fmt.Sprintf("problem while trying to create user storage: %s", err.Error()))
	}

	userService, err := services.NewUserService(userStorage)
	if err != nil {
		return errors.New(fmt.Sprintf("problem while trying to create user service: %s", err.Error()))
	}

	userHandler, err := handlers.NewUserHandler(userService)
	if err != nil {
		return errors.New(fmt.Sprintf("problem while trying to create user handler: %s", err.Error()))
	}

	log.Println("starting http server")
	router := resthttp.NewRouter(userHandler)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return errors.New(fmt.Sprintf("http server starting error: %s", err.Error()))
	}

	return nil
}
