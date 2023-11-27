package web

import (
	"net/http"

	"github.com/brandon-a-pinto/nebula/broker-service/internal/main/factory"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/presentation/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router chi.Router
	Port   string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Port:   port,
		Router: chi.NewRouter(),
	}
}

func (s *WebServer) routes() {
	userHandler := handler.NewUserHandler(
		*factory.CreateUserFactory(),
	)

	s.Router.Post("/users", userHandler.CreateUser)
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	s.routes()
	http.ListenAndServe(s.Port, s.Router)
}
