package app

import (
	"errors"
	"github.com/tanya-lyubimaya/translator/internal/delivery/http"
	"github.com/tanya-lyubimaya/translator/internal/repository/translator"
	"github.com/tanya-lyubimaya/translator/internal/usecase"
	"os"
)

type Application struct {
	server *http.Server
}

func (a *Application) Serve(port string) error {
	return a.server.Serve(port)
}

func New() (*Application, error) {
	translateUrl, apiKey := os.Getenv("TRANSLATEURL"), os.Getenv("APIKEY")
	if translateUrl == "" || apiKey == "" {
		return nil, errors.New("Empty env variables")
	}
	translatorRepo, err := translator.New(translateUrl, apiKey)
	if err != nil {
		return nil, err
	}
	uc, err := usecase.New(translatorRepo)
	if err != nil {
		return nil, err
	}
	server, err := http.New(uc)
	if err != nil {
		return nil, err
	}
	return &Application{server: server}, nil
}
