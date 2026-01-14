package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxirobledo/curso-go/packages/db"
	"github.com/maxirobledo/curso-go/packages/routes"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
