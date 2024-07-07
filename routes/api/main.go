package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"idotno.fr/echo/utils"
)

func ApiRoutes() chi.Router {
	r := chi.NewRouter()

	// ---- Health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// ---- Auth
	r.Post("/auth/register", Register)
	r.Post("/auth/login", Login)

	r.Group(func(r chi.Router) {
		r.Use(utils.JWTAuthMiddleware)

		// ---- Users
		r.Get("/users", ListUsers)
		r.Get("/users/{id:uint}", GetUser)
	})

	return r
}
