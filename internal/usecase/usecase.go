package usecase

import (
	"context"
	"github.com/tanya-lyubimaya/translator/internal/domain"
)

type useCase struct {
	repoTranslator domain.TranslatorRepository
}

func New(repoTranslator domain.TranslatorRepository) (*useCase, error) {
	return &useCase{repoTranslator: repoTranslator}, nil
}

func (uc *useCase) GetTranslation(ctx context.Context, model domain.TranslateRequestModel) (string, error) {
	return uc.repoTranslator.GetTranslation(ctx, model)
}

func (uc *useCase) GetLanguages(ctx context.Context) ([]string, error) {
	return uc.repoTranslator.GetLanguages(ctx)
}

// Close calls all repos Close() methods and does cleanup
func (uc *useCase) Close() {
}
