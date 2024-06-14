package gingonic

import (
	"net/http"

	"github.com/ArjunMalhotra07/apis.git/structs"
	"github.com/gin-gonic/gin"
)

func defaultRoute(c *gin.Context) {
	//! Response Type 1
	// response := gin.H{
	// 	"message": "Hey",
	// }
	//! Response Type 2 using structs
	response := structs.Response{Message: "Hey"}
	c.IndentedJSON(http.StatusOK, response)
}
