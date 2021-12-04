package views

import (
	"net/http"
	"strconv"

	"github.com/mark/educ/models"
)

func Courses(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// username := uadmin.IsAuthenticated(r).User.Username
	context := map[string]interface{}{}

	// Add course
	if r.Method == "POST" {
		dataAction := r.FormValue("data-action")

		if dataAction == "new_course" {
			course := models.Course{}
			course.CourseName = r.FormValue("coursename")
			course.CourseDescription = r.FormValue("coursedesc")
			course.CourseType = r.FormValue("coursetype")
			course.TotalUnits, _ = strconv.Atoi(r.FormValue("courseunits"))
			course.Save()
		}
	}

	return context
}
