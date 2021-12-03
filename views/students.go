package views

import (
	"net/http"

	"github.com/mark/educ/models"
	"github.com/uadmin/uadmin"
)

func Students(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// username := uadmin.IsAuthenticated(r).User.Username
	context := map[string]interface{}{}

	students := []models.Students{}
	uadmin.All(&students)

	context["Students"] = func() interface{} {
		list := []map[string]interface{}{}
		for i := range students {
			uadmin.Preload(&students[i])
			list = append(list, map[string]interface{}{
				"Firstname":  students[i].Firstname,
				"Middlename": students[i].Middlename,
				"Lastname":   students[i].Lastname,
				"Email":      students[i].Email,
				"Cnumber":    students[i].Cnumber,
				"Address":    students[i].Address,
			})
		}
		return list
	}

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
