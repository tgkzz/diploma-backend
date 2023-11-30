package handler

import "net/http"

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.home)
	mux.HandleFunc("/signup", h.singup)
	mux.HandleFunc("/signin", h.signin)
	// mux.HandleFunc("/signout", h.AuthMiddleware(h.signout))

	return h.Handles(mux)
}
