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
		if r.Header.Get("Content-Type") != "application/json" {
			ErrorHandler(w, http.StatusUnsupportedMediaType)
			return
		}

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
		if r.Header.Get("Content-Type") != "application/json" {
			ErrorHandler(w, http.StatusUnsupportedMediaType)
			return
		}

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

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Authorization", token)
		login := fmt.Sprintf("logged in as %s", creds.Login)
		code, err := w.Write([]byte(login))
		if err != nil {
			log.Print(err)
			ErrorHandler(w, code)
			return
		}

		response := map[string]string{"Authorization": token}
		json.NewEncoder(w).Encode(response)

	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) signout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		token := r.Header.Get("Authorization")
		if token == "" {
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}

		w.Write([]byte("logged out"))
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

}
