package controller

import (
	"fmt"
	"net/http"
)

import (
	"github.com/gin-gonic/gin"
)

func SrvHome(c *gin.Context) {
	fmt.Println("SrvHOme")
	c.HTML(http.StatusOK, "index.tpl.html", gin.H{})
	return
}
