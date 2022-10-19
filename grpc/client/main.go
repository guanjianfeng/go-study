package main

import (
	"context"
	"fmt"
	trippb "go_study/grpc/server/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main()  {
	dial, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	client := trippb.NewTripServiceClient(dial)
	trip, err := client.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "123"})
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println(trip)
}
