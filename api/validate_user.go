package api

import (
	"io/ioutil"
	"net/http"

	"github.com/uadmin/uadmin"
	"golang.org/x/crypto/bcrypt"
)

// func ValidateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	userpassword := uadmin.IsAuthenticated(r).User.Password

// 	pass := r.FormValue("password")
// 	Salt, _ := ioutil.ReadFile(".salt")

// 	password := []byte(pass + string(Salt))
// 	hashedPassword := []byte(userpassword)
// 	err_pass := bcrypt.CompareHashAndPassword(hashedPassword, password)

// 	var res *http.Response
// 	var err error

// 	if err_pass == nil {
// 		res, err = http.PostForm("http://0.0.0.0:8080/api/validate_user/",
// 			url.Values{
// 				"password": {pass},
// 			})

// 		fmt.Println("res ", res)

// 		// uadmin.ReturnJSON(w, r, map[string]interface{}{
// 		// 	"msg": "password",
// 		// })

// 	} else {
// 		uadmin.ReturnJSON(w, r, map[string]interface{}{
// 			"msg": "err_password",
// 		})
// 		return
// 	}

// 	if err != nil {
// 		uadmin.ReturnJSON(w, r, map[string]interface{}{
// 			"status":  "error",
// 			"err_msg": "error contacting api. " + err.Error(),
// 		})
// 		return
// 	}
// 	defer res.Body.Close()

// 	buf, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		uadmin.ReturnJSON(w, r, map[string]interface{}{
// 			"status":  "error",
// 			"err_msg": "Error reading response from ValidateUserNameHandler API. " + err.Error(),
// 		})
// 		return
// 	}

// 	obj := map[string]interface{}{}
// 	err = json.Unmarshal(buf, &obj)
// 	fmt.Println("buffff ", buf)
// 	if err != nil {
// 		uadmin.ReturnJSON(w, r, map[string]interface{}{
// 			"status":  "error",
// 			"err_msg": "error parsing json. " + err.Error(),
// 		})
// 	}

// 	if obj["status"] == "ok" {
// 		uadmin.ReturnJSON(w, r, map[string]interface{}{
// 			"msg": "password",
// 		})
// 	}
// 	w.Write(buf)

// }

func ValidateUserHandler(w http.ResponseWriter, r *http.Request) {
	userpassword := uadmin.IsAuthenticated(r).User.Password

	pass := r.FormValue("password")
	Salt, _ := ioutil.ReadFile(".salt")

	password := []byte(pass + string(Salt))
	hashedPassword := []byte(userpassword)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err == nil {

		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"msg": "password",
		})

	} else {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"msg": "err_password",
		})

	}

	// buf :=
	// w.Write()
	// ioutil.ReadAll(r.Body)
}
