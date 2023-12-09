package main

import (
	"fmt"
	"goqdrantvectorsearch/pkg/pack"
	"goqdrantvectorsearch/pkg/qdrantapi"
	"log"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(pack.Hello())
	var connection = qdrantapi.GetConnection()
	qdrantapi.DeleteCollection(connection, "test")
	qdrantapi.CreateCollection(connection, "test", 3)
	fmt.Println(connection)
	var pointsClient = qdrantapi.GetPointsClient(connection)
	qdrantapi.UpsertPoint([]float32{1, 2, 3}, "test", pointsClient)
	var retrievedPoint = qdrantapi.GetPoints(1, "test", pointsClient)
	log.Println(retrievedPoint)
	var searchResult = qdrantapi.Search([]float32{1, -1, 5}, pointsClient, "test")
	log.Println(searchResult)

}
