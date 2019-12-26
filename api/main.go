package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Runner runs api
func Runner() {
	r := gin.Default()
	r.Any("/ping/", controller)

	log.Println("listen on http://0.0.0.0:8080")
	log.Fatal(r.Run("0.0.0.0:8080"))
}
