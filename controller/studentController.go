package controller

import (
	"fmt"
	"net/http"
	"strconv"
)

import (
	"github.com/gin-gonic/gin"
)

import (
	"../model"
	"../service/studentService"
)

// List
func SrvStudents(c *gin.Context) {
	fmt.Println("Srv Student List")
	students := studentService.GetList()
	c.HTML(http.StatusOK, "student.list.tpl.html", gin.H{
		"StudentList": students,
	})
	return
}

// Form
func SrvStudent(c *gin.Context) {
	fmt.Println("Srv Student")
	id := c.Param("id")
	err, student := studentService.Get(id)
	if err != nil {
		fmt.Printf("error: %#v\n", err.Error())
	} else {
		c.HTML(http.StatusOK, "student.form.tpl.html", gin.H{
			"StudentID":   student.StudentID,
			"Name":        student.Name,
			"Description": student.Description,
		})
	}
}

// New
func SrvStudentNewGet(c *gin.Context) {
	fmt.Println("Srv Student New Get")
	c.HTML(http.StatusOK, "student.new.tpl.html", gin.H{})
}
func SrvStudentNew(c *gin.Context) {
	fmt.Println("Srv Student New")

	name := c.PostForm("name")
	description := c.PostForm("description")

	res := studentService.New(model.Student{
		Name:        name,
		Description: description,
	})
	if res == false {
		fmt.Printf("error: SrvStudentNew")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/student/list")
	}
}

// Edit
func SrvStudentEditGet(c *gin.Context) {
	fmt.Println("Srv Student Edit Get")

	id := c.Param("id")
	err, student := studentService.Get(id)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		c.HTML(http.StatusOK, "student.edit.tpl.html", gin.H{
			"StudentID":   student.StudentID,
			"Name":        student.Name,
			"Description": student.Description,
		})
	}
}
func SrvStudentEdit(c *gin.Context) {
	fmt.Println("Srv Student Edit")

	id := c.PostForm("studentID")
	name := c.PostForm("name")
	description := c.PostForm("description")
	intId, err := strconv.Atoi(id)
	if err != nil {
		err := studentService.Update(model.Student{
			StudentID:   intId,
			Name:        name,
			Description: description,
		})
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			c.Redirect(http.StatusMovedPermanently, "/student/list")
		}
	} else {
		fmt.Printf("error: %s\n", err.Error())
	}
}

// Delete
func SrvStudentDelete(c *gin.Context) {
	fmt.Println("Srv Student Delete")
	id := c.Param("id")

	err := studentService.Delete(id)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		c.Redirect(http.StatusMovedPermanently, "/student/list")
	}
}
