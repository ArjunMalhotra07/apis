package gingonic

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GinGonicApiRoutes() {
	router := gin.Default()
	router.GET("/", defaultRoute)
	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
}
