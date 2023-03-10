package main

import (
	"gather-go/package/accounts"
	"gather-go/package/connections/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// viper.SetConfigFile("./package/connections/envs/.env")
	// viper.ReadInConfig()

	// dbURL := viper.Get("DB_URL").(string)
	// port := viper.Get("PORT").(string)
	dbURL := "postgres://steel:sentinel2000@127.0.0.1:5432/gather"
	router := gin.Default()
	dbHandler := db.Init(dbURL)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to gather api"})
	})

	accounts.RegisterRoutes(router, dbHandler)

	router.Run(":8000")
}
