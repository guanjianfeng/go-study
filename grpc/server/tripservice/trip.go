package trip

import (
	"context"
	"fmt"
	trippb "go_study/grpc/server/proto/gen/go"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	trippb.UnimplementedTripServiceServer
}


func (*Service)GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	fmt.Println(c)
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip:&trippb.Trip{
			Start: "abc",
			End: "def",
			DurationSec: 3600,
			FeeCent: 10000,
			StartPos: &trippb.Location{
				Latitude:  1.0,
				Longitude: 2,
			},
			EndPos: &trippb.Location{
				Latitude:  3,
				Longitude: 4,
			},
			PathLocations: []*trippb.Location{
				{Latitude: 5, Longitude: 6},
				{Latitude: 7, Longitude: 8},
			},
			Status: trippb.TripStatus_IN_PROGRESS,
			AddTime: timestamppb.New(time.Now()),
		},
	},nil
}
