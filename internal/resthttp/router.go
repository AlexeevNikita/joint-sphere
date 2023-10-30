package resthttp

import (
	"github.com/AlexeevNikita/joint-sphere/internal/entities"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(
	userHandler entities.UserHandler,
) http.Handler {
	router := chi.NewRouter()

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", userHandler.Create)
		r.Get("/read/{ID}", userHandler.Read)
		r.Post("/update/{ID}", userHandler.Update)
		r.Post("/delete/{ID}", userHandler.Delete)
	})

	return router
}
