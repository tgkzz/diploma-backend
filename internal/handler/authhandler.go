package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	erroring "diploma/internal/model/err"
	"diploma/internal/model/user"
)

func ErrorHandler(w http.ResponseWriter, code int) {
	w.WriteHeader(code)

	fmt.Fprintf(w, "error %d,\n%s", code, http.StatusText(code))
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "hello from home")
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
	}
}

func (h *Handler) singup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user user.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		if err := h.service.AutherService.CreateUser(user); err != nil {
			if err == erroring.ErrInvalidData {
				log.Print(err)
				ErrorHandler(w, http.StatusBadRequest)
				return
			} else if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				log.Print(err)
				ErrorHandler(w, http.StatusConflict)
				return
			} else {
				log.Print(err)
				ErrorHandler(w, http.StatusInternalServerError)
				return
			}
		}

		fmt.Fprintf(w, "success")
	}
}

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var creds user.User

		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		if err := h.service.AutherService.CheckUserCreds(creds); err != nil {
			if err == erroring.ErrIncorrectPassword || strings.Contains(err.Error(), "sql: no rows") {
				log.Print(err)
				ErrorHandler(w, http.StatusUnauthorized)
				return
			} else {
				log.Print(err)
				ErrorHandler(w, http.StatusInternalServerError)
				return
			}
		}

		token, err := h.service.AutherService.CreateToken(creds.Login)
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		w.Header().Add("Token", token)
		login := fmt.Sprintf("logged in as %s", creds.Login)
		code, err := w.Write([]byte(login))
		if err != nil {
			log.Print(err)
			ErrorHandler(w, code)
			return
		}
	}
}
