package handler

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:4000/docs/swagger.json")))

	//auth
	mux.HandleFunc("/", h.home)
	mux.HandleFunc("/signup", h.singup)
	mux.HandleFunc("/signin", h.signin)
	mux.HandleFunc("/signout", h.AuthMiddleware(h.signout))

	// require many improvements
	// payment (test version)
	mux.HandleFunc("/checkout", h.CORSCheck(h.checkout))

	return h.Handles(mux)
}
