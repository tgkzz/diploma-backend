package handler

import (
	"diploma/internal/database"
	"diploma/internal/model"
	"diploma/internal/pkg"
	"diploma/internal/session"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RegistrationhHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user model.User
		var err error
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		if msg, ok := database.CheckUserCreds(user); ok {
			log.Print(msg)
			ErrorHandler(w, http.StatusBadRequest)
			return
		}
		user.Password, err = pkg.HashPassword(user.Password)
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		if err := database.InsertUser(user); err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		code, err := w.Write([]byte("success"))
		if err != nil {
			log.Print(err)
			ErrorHandler(w, code)
			return
		}
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var creds model.User
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		user, err := database.ReturnUser(creds.Login)
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusNotFound)
			return
		}
		if !pkg.CheckPasswordHash(creds.Password, user.Password) {
			log.Print("Invalid Password")
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}
		token, err := session.CreateToken()
		if err != nil {
			log.Print(err)
			log.Println("\nADSAD")
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Token", token)
		login := fmt.Sprintf("logged in as %s", user.Login)
		code, err := w.Write([]byte(login))
		if err != nil {
			log.Print(err)
			ErrorHandler(w, code)
			return
		}
	}
}

// realize the logic of logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pass"))
}
