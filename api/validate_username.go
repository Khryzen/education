package api

import (
	"fmt"
	"net/http"

	"github.com/uadmin/uadmin"
)

func ValidateUserNameHandler(w http.ResponseWriter, r *http.Request) {
	uadmin.Trail(uadmin.DEBUG, "Enter Validate Username")
	user := []uadmin.User{}
	userinput := r.FormValue("newusername")
	uadmin.Trail(uadmin.DEBUG, "userinput: %v", userinput)

	uadmin.Filter(&user, "username=?", userinput)
	fmt.Println("FILTERED: ", user, " len ", len(user))

	if len(user) > 0 {
		fmt.Println("len already exist")
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"validate_username_msg": "username already exist",
		})
		// return

	} else if len(user) == 0 {
		fmt.Println("already exist")
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"validate_username_msg": "username available",
		})
		// return
	}
}
