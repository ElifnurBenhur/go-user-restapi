package main

import (
	"github.com/ElifnurBenhur/go-gin-gorm-try/config"
	"github.com/ElifnurBenhur/go-gin-gorm-try/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8084")
}
