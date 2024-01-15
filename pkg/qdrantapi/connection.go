package qdrantapi

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetConnection() (*grpc.ClientConn, error) {
	var (
		addr = flag.String("addr", "localhost:6334", "the address to connect to")
	)
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	return conn, err
}

// func create_conn() {
