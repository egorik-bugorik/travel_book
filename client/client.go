package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"travel_book/travelbook"
)

func main() {

	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)

	}
	client := travelbook.NewTravelBookClient(conn)

	feature, err := client.GetFeature(context.Background(), &travelbook.Point{
		Latitude:  406421967,
		Longitude: -747727624,
	})
	if err != nil {
		log.Fatalf("Error while get feature for point ::: %v", err)

	}
	log.Println(feature.Name)

}
