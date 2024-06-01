package main

import (
	"gin_api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// n := gin.New()

	router.InitRouter(r)
	r.Run()
}
