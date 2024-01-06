package main

import (
	"fmt"
	"goqdrantvectorsearch/pkg/pack"
	"goqdrantvectorsearch/pkg/qdrantapi"
	"log"
	"net/http"

	// docs "github.com/go-project-name/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	// gin-swagger middleware
)

// swagger embed files

func qdrantExample(c *gin.Context) {
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
	// return c.JSON(200, gin.H{
	// 	"message": "pong",
	// })

}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	r := gin.New()

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	r.Run()
}
