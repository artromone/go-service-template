package v1

import (
	v1 "go-clean-template/internal/controller/amqp_rpc/v1"
	"go-clean-template/internal/usecase"
	"go-clean-template/pkg/logger"
	"go-clean-template/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation, l logger.Interface) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)

	{
		v1.NewTranslationRoutes(routes, t, l)
	}

	return routes
}
