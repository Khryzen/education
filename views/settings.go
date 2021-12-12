package views

import (
	"fmt"
	"net/http"

	"github.com/uadmin/uadmin"
)

func ProfileSettings(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}
	userglobal := r.FormValue("userglobal")
	user := uadmin.User{}
	dataAction := r.FormValue("dataAction")

	if r.Method == "POST" {
		if dataAction == "saveuserchanges" {
			uadmin.Get(&user, "username = ?", userglobal)
			fmt.Println("USER GET", user)
		}
	}

	return context
}
