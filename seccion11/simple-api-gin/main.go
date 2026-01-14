package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status string `json:"string"`
}

func main() {
	// Crea el router con middleware por defecto
	// logger y recovery (panic safe)
	router := gin.Default()

	// Endpoint GET /health
	router.GET("/health", func(c *gin.Context) {
		response := HealthResponse{
			Status: "ok",
		}

		// GIN se encarga del JSON y del status
		c.JSON(http.StatusOK, response)
	})

	// Levantar el servidor
	router.Run(":8080")

}
