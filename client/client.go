package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	"time"
	"travel_book/travelbook"
)

func main() {

	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)

	}
	client := travelbook.NewTravelBookClient(conn)

	featureByPoint(client, &travelbook.Point{
		Latitude:  407838351,
		Longitude: -746143763,
	})

	traversePoint(client)
	featureInsideRectangle(client, &travelbook.Rectangle{
		Lo: &travelbook.Point{
			Latitude:  407838000,
			Longitude: -746143000,
		},
		Hi: &travelbook.Point{
			Latitude:  409838000,
			Longitude: -736143000,
		},
	})

}

func traverseRoute(client travelbook.TravelBookClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("Error while traversing route features with RECT ::: %v", err)

	}
	defer cancel()

}

func featureInsideRectangle(client travelbook.TravelBookClient, rectangle *travelbook.Rectangle) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stream, err := client.ListFeature(ctx, rectangle)
	if err != nil {
		log.Fatalf("Error while getting features with RECT ::: %v", err)

	}
gori:
	for {
		feature, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break gori
			}
			log.Fatalf("Error while getting features with RECT ::: %v", err)

		}
		log.Printf("\n\n\t :::   Feature found -> %v   :::", feature.Name)
	}

}

func featureByPoint(client travelbook.TravelBookClient, p *travelbook.Point) {

	ctx := context.Background()
	feature, err := client.GetFeature(ctx, p)
	if err != nil {
		log.Fatalf("Error while getting feature by point!!! %v", err)
	}
	log.Printf("\n:::    Feature  :::    %v", feature.Name)
}

func randomPoint() *travelbook.Point {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	lat := (r.Intn(180) - 90) * 1e7
	lon := (r.Intn(360) - 180) * 1e7
	return &travelbook.Point{
		Latitude:  int32(lat),
		Longitude: int32(lon),
	}
}
