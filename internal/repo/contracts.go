// Package repo implements application outer layer logic. Each logic group in own file.
package repo

import (
	"context"
	"go-clean-template/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_repo_test.go -package=usecase_test

type TranslationRepo interface {
	Store(context.Context, entity.Translation) error
	GetHistory(context.Context) ([]entity.Translation, error)
}

type TranslationWebAPI interface {
	Translate(entity.Translation) (entity.Translation, error)
}
