package views

import (
	"net/http"

	"github.com/mark/educ/models"
)

func Students(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// username := uadmin.IsAuthenticated(r).User.Username

	context := map[string]interface{}{}
	students := []models.Students{}
	// Filter the tags that matches the created by to the username of the active user
	// uadmin.Filter(&tags, "created_by = ?", username)
	// Assign the values of tags into the context
	context["students"] = students

	return context
}
