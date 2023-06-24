package main

import (
	"github.com/imvahid/go-crud/initializers"
	"github.com/imvahid/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
