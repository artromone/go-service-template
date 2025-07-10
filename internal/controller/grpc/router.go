package grpc

import (
	v1 "go-clean-template/internal/controller/grpc/v1"
	"go-clean-template/internal/usecase"
	"go-clean-template/pkg/logger"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewRouter -.
func NewRouter(app *pbgrpc.Server, t usecase.Translation, l logger.Interface) {
	{
		v1.NewTranslationRoutes(app, t, l)
	}

	reflection.Register(app)
}
