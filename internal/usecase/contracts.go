// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"go-clean-template/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

// Translation -.
type Translation interface {
	Translate(context.Context, entity.Translation) (entity.Translation, error)
	History(context.Context) (entity.TranslationHistory, error)
}
