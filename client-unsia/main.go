package main

import (
	"client-unsia/pb/cities"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %s", err)
		return
	}
	defer conn.Close()

	ctx := context.Background()

	city := cities.NewCitiesServiceClient(conn)

	response, err := city.GetCity(ctx, &cities.City{
		Id: 1,
	})

	if err != nil {
		fmt.Printf("Error when calling grpc service: %s", err)
		return
	}
	fmt.Printf("Resp received: %v", response)
}
