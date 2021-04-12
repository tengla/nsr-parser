package trainroute

import (
	"context"
	"fmt"

	pb "github.com/tengla/nsr-parser/trainroute/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Foo() string {
	return "Bar"
}

type OtrServer struct {
	pb.UnimplementedOperatingTrainRouteServiceServer
}

func NewOtrServer() *OtrServer {
	return &OtrServer{}
}

// GetItem - rpc impl
func (s *OtrServer) GetItems(ctx context.Context, ids *pb.OtrIdList) (*pb.OperatingTrainRouteList, error) {
	list := createList()
	found := &pb.OperatingTrainRouteList{}
	for _, oid := range list.TrainIds {
		for _, id := range ids.TrainIds {
			if oid.TrainId == id.TrainId {
				found.Routes = append(found.Routes, &pb.OperatingTrainRoute{
					TrainId:   oid.TrainId,
					TrainDate: timestamppb.Now(),
				})
			}
		}
	}
	// status.Errorf(codes.NotFound, "route %s not found", id.TrainId)
	return found, nil
}

func createList() *pb.OtrIdList {
	list := &pb.OtrIdList{}
	ids := make([]int, 500)
	list.TrainIds = make([]*pb.OtrId, 0)
	for id := range ids {
		train_id := fmt.Sprintf("%.3d", id+1)
		otrId := &pb.OtrId{
			TrainId: train_id,
		}
		list.TrainIds = append(list.TrainIds, otrId)
	}
	return list
}

// Fullstate - rpc impl
func (s *OtrServer) FullState(ctx context.Context, e *pb.Empty) (*pb.OtrIdList, error) {
	return createList(), nil
}
