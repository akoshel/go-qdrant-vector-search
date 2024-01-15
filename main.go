package main

import (
	_ "goqdrantvectorsearch/docs"
	routers "goqdrantvectorsearch/routers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// New gin router
	router := gin.New()

	// Register api/v1 endpoints
	routers.Register(router)
	router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))

	// Listen and Server in
	router.Run()
}
