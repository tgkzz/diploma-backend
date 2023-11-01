package handler

import (
	"diploma/internal/database"
	"diploma/internal/model"
	"diploma/internal/pkg"
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

// TODO : generate JWT Tokens, check auth on every handler which requires it, logout from system
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
		login := fmt.Sprintf("logged in as %s", user.Login)
		code, err := w.Write([]byte(login))
		if err != nil {
			log.Print(err)
			ErrorHandler(w, code)
			return
		}
	}
}
