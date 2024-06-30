package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"idotno.fr/echo/routes/api"
	"idotno.fr/echo/services"
	"idotno.fr/echo/utils"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// ---- Frontend
	if utils.GetEnv("ECHO_ENABLE_FRONTEND", "true") == "true" {
		r.Get("/", RenderTemplate("templates/home.gohtml"))
		r.Get("/auth", RenderTemplate("templates/auth.gohtml"))

		r.Group(func(r chi.Router) {
			r.Use(utils.JWTAuthMiddleware)

			// ---- Users
			r.Get("/users", RenderTemplate("templates/users.gohtml"))

			// ---- Chat
			r.Get("/chat", RenderTemplate("templates/chat.gohtml"))
		})
	}

	// ---- WS
	r.With(utils.JWTAuthMiddleware).HandleFunc("/ws", services.HandleWsConnections)

	// ---- API
	r.Mount("/api", api.ApiRoutes())

	return r
}
