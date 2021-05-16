module kratos-http-grpc

go 1.15

require (
	github.com/go-kratos/kratos/v2 v2.0.0-beta4
	github.com/google/wire v0.5.0
	github.com/gorilla/mux v1.8.0
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.20.0
	go.opentelemetry.io/otel/sdk v0.20.0
	go.opentelemetry.io/otel/trace v0.20.0
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.2.2
)
