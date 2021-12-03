package views

import (
	"net/http"

	"github.com/mark/educ/models"
)

func Students(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// username := uadmin.IsAuthenticated(r).User.Username
	context := map[string]interface{}{}

	// Add student
	if r.Method == "POST" {
		dataAction := r.FormValue("data-action")

		if dataAction == "new_student" {
			student := models.Students{}
			student.Firstname = r.FormValue("firstname")
			student.Middlename = r.FormValue("middlename")
			student.Lastname = r.FormValue("lastname")
			student.Email = r.FormValue("email")
			student.Cnumber = r.FormValue("cnumber")
			student.Address = r.FormValue("address")

			student.Save()
		}
	}

	return context
}
