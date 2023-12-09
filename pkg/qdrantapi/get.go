package qdrantapi

import (
	"context"
	"log"
	"time"

	pb "github.com/qdrant/go-client/qdrant"
)

func GetPoints(id uint64, collectionName string, pointsClient pb.PointsClient) []*pb.RetrievedPoint {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	pointsById, err := pointsClient.Get(ctx, &pb.GetPoints{
		CollectionName: collectionName,
		Ids: []*pb.PointId{
			{PointIdOptions: &pb.PointId_Num{Num: id}},
		},
	})
	if err != nil {
		log.Fatalf("Could not retrieve points: %v", err)
	}
	return pointsById.GetResult()
}
