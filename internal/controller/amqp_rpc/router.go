package amqprpc

import (
	"github.com/ZNick65/go-clean-template/internal/usecase"
	"github.com/ZNick65/go-clean-template/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
