package procedures

import (
	"context"

	"github.com/eser/acik.io/pkg/bliss/di"
	"github.com/eser/acik.io/pkg/bliss/grpcfx"
	pb "github.com/eser/acik.io/pkg/proto/broadcast"
)

type BroadcastService struct {
	pb.UnimplementedChannelServiceServer
	pb.UnimplementedMessageServiceServer
}

func RegisterBroadcastService(container di.Container, grpcService grpcfx.GrpcService) {
	bs := NewBroadcastService()

	grpcService.RegisterService(&pb.ChannelService_ServiceDesc, bs)
	grpcService.RegisterService(&pb.MessageService_ServiceDesc, bs)
}

func NewBroadcastService() *BroadcastService {
	return &BroadcastService{} //nolint:exhaustruct
}

func (s *BroadcastService) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Channel, error) {
	channel := &pb.Channel{
		Id:   "123",
		Name: "Test Channel",
	}

	return channel, nil
}

func (s *BroadcastService) List(ctx context.Context, req *pb.ListRequest) (*pb.Channels, error) {
	// Implementation here
	return nil, nil //nolint:nilnil
}

func (s *BroadcastService) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	// Implementation here
	return nil, nil //nolint:nilnil
}
