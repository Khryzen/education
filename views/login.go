package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	context := map[string]interface{}{}
	session := uadmin.IsAuthenticated(r)

	if session != nil {
		http.Redirect(w, r, "/educ/dashboard/", http.StatusSeeOther)
		return
	}

	user := uadmin.User{}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			context["login_denied"] = true
			context["err_msg"] = "Please input your username."
		}

		if password == "" {
			context["login_denied"] = true
			context["err_msg"] = "Please input your password."
		}

		if username == "" && password == "" {
			context["login_denied"] = true
			context["err_msg"] = "Please input your username and password."
		}

		if username != "" && password != "" {
			uadmin.Get(&user, "username=?", username)

			session := user.Login(password, "")

			if session != nil {
				http.SetCookie(w, &http.Cookie{
					Path:     "/",
					Name:     "session",
					Value:    session.Key,
					SameSite: http.SameSiteStrictMode,
				})

				http.Redirect(w, r, "/educ/dashboard/", 303)
				return
			} else {
				context["login_denied"] = true
				context["err_msg"] = "Login denied"
			}
		}
	}
	uadmin.RenderHTML(w, r, "./templates/login.html", context)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	session := uadmin.Session{}
	key := uadmin.GetUserFromRequest(r).GetActiveSession().Key
	uadmin.Get(&session, "`key` = ? ", key)
	uadmin.DeleteList(&session, "`key` = ?", key)

	// Expire session cookie on logout
	sessionCookie := &http.Cookie{
		Name:   "session",
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, sessionCookie)
	http.Redirect(w, r, "/login/", http.StatusSeeOther)
}
