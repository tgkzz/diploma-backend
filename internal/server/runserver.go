package server

import (
	"diploma/internal/handler"
	"diploma/internal/session"
	"net/http"
)

func MakeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", handler.RegistrationhHandler)
	mux.HandleFunc("/login", handler.LoginHandler)
	mux.Handle("/logout", session.CheckAuth(handler.LogoutHandler))

	return mux
}
