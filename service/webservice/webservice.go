package webservice

import (
	"fmt"
	"net/http"
)

import (
	"github.com/gin-gonic/gin"
)

import (
	"../../controller"
)

func StartWebServer() {
	fmt.Println("Start WebServer")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Home
	r.GET("/", controller.SrvHome)

	// Student
	student := r.Group("/student")
	{
		student.GET("/list", controller.SrvStudents)
		student.GET("/new", controller.SrvStudentNewGet)
		student.POST("/new", controller.SrvStudentNew)
		student.GET("/form/:id", controller.SrvStudent)
		student.GET("/form/:id/edit", controller.SrvStudentEditGet)
		student.POST("/form/:id/edit", controller.SrvStudentEdit)
		student.DELETE("/form/:id/delete", controller.SrvStudentDelete)
	}

	// Others
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}
