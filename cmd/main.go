package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Small list of quotes
var quotes = []string{
	"Believe in yourself!",
	"Keep pushing forward.",
	"Every day is a new opportunity.",
	"Stay positive, work hard, make it happen.",
	"Success is not final, failure is not fatal.",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()

	router.GET("/quote", func(c *gin.Context) {
		quote := quotes[rand.Intn(len(quotes))]
		c.JSON(http.StatusOK, gin.H{"quote": quote})
	})

	router.Run(":8080") // Run on http://localhost:8080
}
