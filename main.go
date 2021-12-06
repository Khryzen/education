package main

import (
	"net/http"

	"github.com/mark/educ/models"
	"github.com/mark/educ/views"
	"github.com/uadmin/uadmin"
)

func main() {
	DBConfig()
	RegisterModels()
	Handlers()
	PageSetup()
}

func DBConfig() {
	uadmin.Database = &uadmin.DBSettings{
		Type:     "mysql",
		Host:     "localhost",
		Name:     "educ",
		User:     "root",
		Password: "Allen is Great 200%",
		Port:     3306,
	}
}

func RegisterModels() {
	uadmin.Register(
		models.Students{},
		models.Course{},
		models.Subjects{},
	)
}

func Handlers() {
	http.HandleFunc("/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	http.HandleFunc("/educ/", uadmin.Handler(views.EducHandler))
}

func PageSetup() {
	uadmin.RootURL = "/admin/"
	uadmin.Port = 8080
	uadmin.SiteName = "Education System"
	uadmin.StartServer()
}
