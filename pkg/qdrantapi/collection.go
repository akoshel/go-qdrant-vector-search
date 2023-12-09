package qdrantapi

import (
	"context"
	"log"
	"time"

	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
)

func CreateCollection(conn *grpc.ClientConn, collectionName string, vectorSize uint64) {
	collections_client := pb.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	var defaultSegmentNumber uint64 = 2
	// Create collection
	_, err := collections_client.Create(ctx, &pb.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     vectorSize,
				Distance: pb.Distance_Cosine,
			},
		}},
		OptimizersConfig: &pb.OptimizersConfigDiff{
			DefaultSegmentNumber: &defaultSegmentNumber,
		},
	})
	if err != nil {
		log.Fatalf("Could not create collection: %v", err)
	} else {
		log.Println("Collection", collectionName, "created")
	}
}

func DeleteCollection(conn *grpc.ClientConn, collectionName string) {
	collections_client := pb.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err := collections_client.Delete(ctx, &pb.DeleteCollection{
		CollectionName: collectionName,
	})
	if err != nil {
		log.Fatalf("Could not delete collection: %v", err)
	} else {
		log.Println("Collection", collectionName, "deleted")
	}
}
