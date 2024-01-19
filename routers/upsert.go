package routers

import (
	// "goqdrantvectorsearch/pkg/qdrantapi"

	"goqdrantvectorsearch/pkg/qdrantapi"

	"github.com/gin-gonic/gin"
)

type NewVector struct {
	Data           []float32
	CollectionName string
}

// @Summary Upsert vector
// @Description post string by ID
// @Accept  json
// @Produce  json
// @Success 200 "ok"
// @Failure 503 "error"
// @Param newVector body NewVector true "New vector"
// @Router /upsertVector [post]
func UpsertVector(c *gin.Context) {
	var newVector NewVector
	err := c.ShouldBind(&newVector)
	if err != nil {
		c.JSON(503, err)
		return
	}
	connection, conn_err := qdrantapi.GetConnection()
	if conn_err != nil {
		c.JSON(503, err)
		return
	}
	client := qdrantapi.GetPointsClient(connection)
	upsert_err := qdrantapi.UpsertPoint(newVector.Data, newVector.CollectionName, client)
	if upsert_err != nil {
		c.JSON(503, upsert_err)
	} else {
		c.JSON(200, "success")
	}

}
