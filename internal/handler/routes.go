package handler

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:4000/docs/swagger.json")))

	//auth
	mux.HandleFunc("/api", h.home)
	mux.HandleFunc("/api/register", h.register)
	mux.HandleFunc("/api/login", h.login)
	mux.HandleFunc("/api/logout", h.AuthMiddleware(h.logout))

	// require many improvements
	// payment (test version)
	// mux.HandleFunc("/checkout", h.CORSCheck(h.checkout))

	return h.Handles(mux)
}
