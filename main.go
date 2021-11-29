package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var amountOfHits = 0

func main() {
	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		// The problem is here, Golang does multi-threading by default, they're called Goroutines
		// Here is a hint ;) https://go.dev/tour/concurrency/9
		amountOfHits++
		context.String(200, fmt.Sprintf("Amount of hits: %d", amountOfHits))
	})

	server.GET("/reset", func(context *gin.Context) {
		amountOfHits = 0
		context.String(200, fmt.Sprintf("Amount of hits: %d", amountOfHits))
	})

	server.Run(":1997")
}
