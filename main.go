package main

import (
	"goapiauth/db"
	"goapiauth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
