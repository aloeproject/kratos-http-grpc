package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/trace"
	v1 "kratos-http-grpc/api/vehicle/v1"
	"kratos-http-grpc/app/vehicle/internal/conf"
	"kratos-http-grpc/app/vehicle/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, tracer trace.TracerProvider, vehicle *service.VehicleService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	m := http.Middleware(
		middleware.Chain(
			tracing.Server(tracing.WithTracerProvider(tracer)),
			recovery.Recovery(),
			logging.Server(logger),
		),
	)

	srv.HandlePrefix("/", v1.NewVehicleHandler(vehicle, m))
	return srv
}
