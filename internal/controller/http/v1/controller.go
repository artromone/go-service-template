package v1

import (
	"go-clean-template/internal/usecase"
	"go-clean-template/pkg/logger"

	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}

// TODO: for what
