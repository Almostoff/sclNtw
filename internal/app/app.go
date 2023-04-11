package app

import (
	"github.com/go-chi/chi/v5"
	"lesson30/internal/cases"
	"lesson30/internal/controller"
	"lesson30/internal/repository"
	"net/http"
)

func Run() {
	repository := repository.NewRepository()
	usecase := cases.NewUseCase(repository)
	r := chi.NewRouter()
	controller.Build(r, usecase)
	http.ListenAndServe("localhost:8080", r)
}
