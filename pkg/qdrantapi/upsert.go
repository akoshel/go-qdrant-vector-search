package qdrantapi

import (
	"context"
	"log"
	"time"

	pb "github.com/qdrant/go-client/qdrant"
)

func UpsertPoint(data []float32, collectionName string, pointsClient pb.PointsClient) {
	points := []*pb.PointStruct{
		{
			// Point Id is number or UUID
			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Num{Num: 1},
			},
			Vectors: &pb.Vectors{VectorsOptions: &pb.Vectors_Vector{Vector: &pb.Vector{Data: data}}},
			Payload: map[string]*pb.Value{
				"city": {
					Kind: &pb.Value_StringValue{StringValue: "Berlin"},
				},
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	waitUpsert := true
	_, err := pointsClient.Upsert(ctx, &pb.UpsertPoints{
		CollectionName: collectionName,
		Wait:           &waitUpsert,
		Points:         points,
	})
	if err != nil {
		log.Fatalf("Could not upsert points: %v", err)
	} else {
		log.Println("Upsert", len(points), "points")
	}
}
