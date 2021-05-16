package service

import (
	"context"
	"strconv"

	pb "kratos-http-grpc/api/vehicle/v1"
)

type VehicleService struct {
	pb.UnimplementedVehicleServer
}

func NewVehicleService() *VehicleService {
	return &VehicleService{}
}

func (s *VehicleService) GetVehicle(ctx context.Context, req *pb.GetVehicleRequest) (*pb.GetVehicleReply, error) {
	var res pb.GetVehicleReply
	res.Name = "汽车:" + strconv.FormatInt(req.Id, 10)

	return &res, nil
}
