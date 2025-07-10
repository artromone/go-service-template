package v1

import (
	v1 "go-clean-template/docs/proto/v1"
	"go-clean-template/internal/usecase"
	"go-clean-template/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	v1.TranslationServer

	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}
