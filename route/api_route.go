package route

import (
	"github.com/KaySar12/NextZen-MessageBus/codegen"
	"github.com/KaySar12/NextZen-MessageBus/service"
	jsoniter "github.com/json-iterator/go"
)

type APIRoute struct {
	services *service.Services
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewAPIRoute(services *service.Services) codegen.ServerInterface {
	return &APIRoute{
		services: services,
	}
}
