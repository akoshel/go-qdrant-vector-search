package qdrantapi

import (
	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
)

func GetPointsClient(conn *grpc.ClientConn) pb.PointsClient {
	pointsClient := pb.NewPointsClient(conn)
	return pointsClient
}
