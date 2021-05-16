package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.opentelemetry.io/otel/trace"
	v1 "kratos-http-grpc/api/vehicle/v1"
	"strconv"
	"time"

	pb "kratos-http-grpc/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	tracer trace.TracerProvider
}

func NewUserService(tracer trace.TracerProvider) *UserService {
	return &UserService{tracer: tracer}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	var res pb.GetUserReply

	res.Name = strconv.FormatInt(req.Id, 10)
	conn, err := grpc.DialInsecure(ctx,
		grpc.WithEndpoint("127.0.0.1:9012"),
		grpc.WithMiddleware(middleware.Chain(
			tracing.Server(tracing.WithTracerProvider(s.tracer)),
			recovery.Recovery())),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		return nil, err
	}

	vObj := v1.NewVehicleClient(conn)
	reply, err := vObj.GetVehicle(ctx, &v1.GetVehicleRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	res.Name = reply.Name
	return &res, nil
}
