package main

import (
	"github.com/yusril86/gin-rest-api/initializer"
	"github.com/yusril86/gin-rest-api/models"
)

func init() {
	initializer.LoadEnvVariablesss()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.Post{})
}
