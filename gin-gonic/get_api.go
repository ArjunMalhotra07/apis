package gingonic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApi() {
	router := gin.Default()
	router.GET("/", defaultRoute)
	err := router.Run("8080")
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
}

func defaultRoute(c *gin.Context) {
	response := gin.H{
		"message": "Hey",
	}
	c.IndentedJSON(http.StatusOK, response)
}
