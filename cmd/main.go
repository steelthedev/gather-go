package main

import (
	"gather-go/package/accounts"
	"gather-go/package/connections/db"
	"gather-go/package/room"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {

	dbURL := "postgres://steel:password@127.0.0.1:5432/gather"
	router := gin.Default()
	dbHandler := db.Init(dbURL)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to gather api"})
	})

	router.Use(func(c *gin.Context) {
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		}).ServeHTTP(c.Writer, c.Request, func(w http.ResponseWriter, r *http.Request) {
		})
	})

	accounts.RegisterRoutes(router, dbHandler)
	room.RegisterRoutes(router, dbHandler)

	router.Run(":8000")
}
