package handler

import (
	"net/http"
)

func (h *Handler) Handles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" || !h.service.AutherService.ValidateToken(token) {
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
