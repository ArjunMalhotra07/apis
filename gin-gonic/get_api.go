package gingonic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"greeting"`
}

func GinGonicApiRoutes() {
	router := gin.Default()
	router.GET("/", defaultRoute)
	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
}

func defaultRoute(c *gin.Context) {
	//! Response Type 1
	// response := gin.H{
	// 	"message": "Hey",
	// }
	//! Response Type 2
	response := Response{Message: "Hey"}
	c.IndentedJSON(http.StatusOK, response)
}
