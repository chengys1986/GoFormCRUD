package service

import (
	"bytes"
	"fmt"
	"html/template"
)

import (
	"../model"
)

var (
	HeaderTpl string = "template/shared/header.tpl.html"
	NavbarTpl string = "template/shared/navbar.tpl.html"
	FooterTpl string = "template/shared/footer.tpl.html"
)

func GetHomePage() string {
	fmt.Println("service: GetHomePage...")
	var docHTML bytes.Buffer
	loadTemplate("template/index.html").ExecuteTemplate(&docHTML, "indexPage", nil)
	return docHTML.String()
}

func GetStudentListPage(students []model.Student) string {
	fmt.Println("service: GetStudentListPage...")

	var tplParam = struct {
		Type        string
		StudentList []model.Student
	}{
		"LIST",
		students,
	}

	var docHTML bytes.Buffer
	err := loadTemplate("template/student/student.list.tpl.html").ExecuteTemplate(&docHTML, "studentListPage", tplParam)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	return docHTML.String()
}

func loadTemplate(tplPath string) *template.Template {
	return template.Must(template.ParseFiles(HeaderTpl, NavbarTpl, FooterTpl, tplPath))
}
