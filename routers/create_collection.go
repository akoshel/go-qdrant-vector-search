package routers

import (
	// "goqdrantvectorsearch/pkg/qdrantapi"

	"fmt"

	"goqdrantvectorsearch/pkg/qdrantapi"

	"github.com/gin-gonic/gin"
)

type NewCollection struct {
	Name       string
	VectorSize uint64
}

type ToRemoveCollection struct {
	Name string
}

// @Summary Creates Qdrant Collection
// @Description post string by ID
// @Accept  json
// @Produce  json
// @Success 200 "ok"
// @Failure 503 "error"
// @Param newCollection body NewCollection true "New collection description"
// @Router /createCollection [post]
func CreateCollection(c *gin.Context) {
	var newCollection NewCollection
	err := c.ShouldBind(&newCollection)
	if err == nil {
		fmt.Println(newCollection)
		connection, connection_error := qdrantapi.GetConnection()
		if connection_error != nil {
			c.JSON(503, "Connection error")
			return
		}
		defer connection.Close()
		create_collection_err := qdrantapi.CreateCollection(connection, newCollection.Name, newCollection.VectorSize)
		if create_collection_err == nil {
			c.JSON(200, "")
		} else {
			fmt.Println(create_collection_err)
			c.JSON(503, "Qdrant collection error")
		}
	} else {
		fmt.Println(err)
		c.JSON(503, "Incorrect structure")
	}
}

// @Summary Creates Qdrant Collection
// @Description post string by ID
// @Accept  json
// @Produce  json
// @Success 200 "ok"
// @Failure 503 "error"
// @Param removeCollection body ToRemoveCollection true "Remove collection description"
// @Router /removeCollection [post]
func RemoveCollection(c *gin.Context) {
	var removeCollection ToRemoveCollection
	err := c.ShouldBind(&removeCollection)
	if err == nil {
		fmt.Println(removeCollection)
		connection, connection_error := qdrantapi.GetConnection()

		if connection_error != nil {
			c.JSON(503, connection_error)
			return
		}
		defer connection.Close()
		delete_connection_err := qdrantapi.DeleteCollection(connection, removeCollection.Name)
		if delete_connection_err != nil {
			c.JSON(503, delete_connection_err)
			return
		}
		c.JSON(200, removeCollection.Name)

	} else {
		c.JSON(503, err)
	}
}
