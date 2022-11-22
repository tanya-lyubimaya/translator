package domain

import (
	"context"
)

type TranslatorRepository interface {
	GetLanguages(ctx context.Context) ([]string, error)
	GetTranslation(ctx context.Context, input TranslateRequestModel) (string, error)
}

type UseCase interface {
	Close()
	GetTranslation(ctx context.Context, input TranslateRequestModel) (string, error)
	GetLanguages(ctx context.Context) ([]string, error)
}
