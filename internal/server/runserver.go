package server

import (
	"diploma/internal/handler"
	"net/http"
)

func MakeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", handler.RegistrationhHandler)
	mux.HandleFunc("/login", handler.LoginHandler)

	return mux
}
