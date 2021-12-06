package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func EducHandler(w http.ResponseWriter, r *http.Request) {
	session := r.FormValue("session")
	if session != "" {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: session,
			Path:  "/",
		})
	}

	sess := uadmin.IsAuthenticated(r)
	if sess == nil {
		http.Redirect(w, r, "/", 303)
		return
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/educ/")
	page := strings.TrimSuffix(r.URL.Path, "/")
	context := map[string]interface{}{}

	switch page {
	case "dashboard":
		context = Dashboard(w, r)
	case "students":
		context = Students(w, r)
	case "courses":
		context = Courses(w, r)
	case "registrar":
		context = Registrar(w, r)
	default:
		page = "dashboard"
		context = Dashboard(w, r)
	}

	context["Page"] = strings.Title(page)
	EducRender(w, r, page, context)
}
func EducRender(w http.ResponseWriter, r *http.Request, tpl string, context map[string]interface{}) {
	// Setting the username
	var username string
	if uadmin.GetUserFromRequest(r).String() != "" {
		username = uadmin.GetUserFromRequest(r).String()
	} else {
		username = uadmin.GetUserFromRequest(r).Username
	}

	context["Username"] = username

	templateList := []string{}
	templateList = append(templateList, "./templates/educ/base.html")

	path := "./templates/educ/" + tpl + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
