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

	traverseRoute(client)
	//featureInsideRectangle(client, &travelbook.Rectangle{
	//	Lo: &travelbook.Point{
	//		Latitude:  407838000,
	//		Longitude: -746143000,
	//	},
	//	Hi: &travelbook.Point{
	//		Latitude:  409838000,
	//		Longitude: -736143000,
	//	},
	//})

}

func traverseRoute(client travelbook.TravelBookClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	defer cancel()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	stream, err := client.RecordRoute(ctx)
	var points []*travelbook.Point
	for i := 0; i < r.Intn(100)+2; i++ {
		points = append(points, randomPoint())
	}
	if err != nil {
		log.Fatalf("Error while connecting to traverseroute grpc features  ::: %v", err)

	}
	for _, point := range points {
		point = randomPoint()
		err := stream.Send(point)
		if err != nil {
			log.Fatalf("Error while travers point ::: %v", err)

		}

		log.Printf("\n%d points has been registred successfully!!!", len(points))
	}

	println("end of sending!!!!")
	time.Sleep(time.Second * 5)
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while recieving routeSummary ::: %v", err)

	}
	log.Printf("\n \t::: RESULT SUMMARY :::\n \t %v", res)

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
