package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"webapp/docs"
)

var router *gin.Engine

func main() {
	// set gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// set the router as the default one provided by gin
	router = gin.Default()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Web Application Practice"
	docs.SwaggerInfo.Description = "This is an Api Documentation for my Web Application practice."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// initialize the routes
	initRoute()

	// start serving the app
	router.Run(":8080")
}
