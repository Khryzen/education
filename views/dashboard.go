package views

import (
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// username := uadmin.IsAuthenticated(r).User.Username

	context := map[string]interface{}{}
	// tags := []models.Tag{}
	// Filter the tags that matches the created by to the username of the active user
	// uadmin.Filter(&tags, "created_by = ?", username)
	// Assign the values of tags into the context
	// context["tags"] = tags

	return context
}
