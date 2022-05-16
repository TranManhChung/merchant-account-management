package infra

import (
	"github.com/gorilla/mux"
	"log"
	"main.go/common/err"
	"net/http"
)

type App struct {
	router *mux.Router
}

func New(router *mux.Router) (*App, error) {
	if router == nil {
		return nil, err.NilRouter
	}
	return &App{
		router: router,
	}, nil
}

func (a *App) Start() error {
	log.Printf("Starting server at port %d", 8080)
	return http.ListenAndServe(":8080", a.router)
}
