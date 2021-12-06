package views

import (
	"html/template"
	"net/http"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session := uadmin.IsAuthenticated(r)

	if session != nil {
		http.Redirect(w, r, "/educ/dashboard/", http.StatusSeeOther)
		return
	}

	user := uadmin.User{}

	if r.Method == "POST" {
		if r.FormValue("request") == "Log In" {
			username := r.FormValue("username")
			password := r.FormValue("password")

			uadmin.Trail(uadmin.DEBUG, "Username: %v", username)
			uadmin.Trail(uadmin.DEBUG, "Password: %v", password)

			uadmin.Get(&user, "username=?", username)
			uadmin.Trail(uadmin.DEBUG, "USERNAME: %v", user.Username)
			uadmin.Trail(uadmin.DEBUG, "PASSWORD: %v", user.Password)
			session := user.Login(password, "")

			if session != nil {
				http.SetCookie(w, &http.Cookie{
					Path:     "/",
					Name:     "session",
					Value:    session.Key,
					SameSite: http.SameSiteStrictMode,
				})
				http.Redirect(w, r, "/educ/dashboard/", http.StatusSeeOther)
				return
			}

		}
	}
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		uadmin.Trail(uadmin.ERROR, "Error %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "login.html", "")
	uadmin.Trail(uadmin.DEBUG, "ERR: %v", err)

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
