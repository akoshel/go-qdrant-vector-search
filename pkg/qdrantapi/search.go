package qdrantapi

import (
	"context"
	"log"
	"time"

	pb "github.com/qdrant/go-client/qdrant"
)

func Search(data []float32, pointsClient pb.PointsClient, collectionName string) *pb.SearchResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	unfilteredSearchResult, err := pointsClient.Search(ctx, &pb.SearchPoints{
		CollectionName: collectionName,
		Vector:         data,
		Limit:          3,
		// Include all payload and vectors in the search result
		WithVectors: &pb.WithVectorsSelector{SelectorOptions: &pb.WithVectorsSelector_Enable{Enable: true}},
		WithPayload: &pb.WithPayloadSelector{SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true}},
	})
	if err != nil {
		log.Fatalf("Could not search points: %v", err)
	} else {
		log.Printf("Found points: %s", unfilteredSearchResult.GetResult())
	}
	return unfilteredSearchResult
}
