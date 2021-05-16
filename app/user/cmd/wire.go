// The build tag makes sure the stub is not built in the final build.
//+build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"kratos-http-grpc/app/user/internal/conf"
	"kratos-http-grpc/app/user/internal/server"
	"kratos-http-grpc/app/user/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
