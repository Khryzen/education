package views

import (
	"net/http"

	"github.com/mark/educ/models"
)

func Dashboard(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// username := uadmin.IsAuthenticated(r).User.Username

	context := map[string]interface{}{}
	// Add Course
	if r.Method == "POST" {
		dataAction := r.FormValue("data-action")

		if dataAction == "new_college" {
			college := models.College{}
			college.CollegeName = r.FormValue("college")
			college.CollegeDescription = r.FormValue("collegedesc")
			college.FieldMajor = r.FormValue("fieledmajor")
			college.CollegeEmail = r.FormValue("collegeemail")
			college.CollegeNumber = r.FormValue("collegenumber")

			college.Save()
		}
	}
	return context
}
