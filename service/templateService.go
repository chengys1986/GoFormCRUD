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

func GetStudentNewPage() string {
	fmt.Println("service: GetStudentNewPage...")

	var docHTML bytes.Buffer
	err := loadTemplate("template/student/student.new.tpl.html").ExecuteTemplate(&docHTML, "studentNewPage", nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	return docHTML.String()
}

func GetStudentFormPage(student model.Student) string {
	fmt.Println("service: GetStudentFormPage...")
	fmt.Printf("%#v\n", student)
	var docHTML bytes.Buffer
	err := loadTemplate("template/student/student.form.tpl.html").ExecuteTemplate(&docHTML, "studentFormPage", student)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	return docHTML.String()
}

func GetStudentEditPage(student model.Student) string {
	fmt.Println("service: GetStudentEditPage...")
	fmt.Printf("%#v\n", student)
	var docHTML bytes.Buffer
	err := loadTemplate("template/student/student.edit.tpl.html").ExecuteTemplate(&docHTML, "studentEditPage", student)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	return docHTML.String()
}
func loadTemplate(tplPath string) *template.Template {
	return template.Must(template.ParseFiles(HeaderTpl, NavbarTpl, FooterTpl, tplPath))
}
