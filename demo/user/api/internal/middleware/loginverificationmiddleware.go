package middleware

import "net/http"

type LoginVerificationMiddleware struct {
}

func NewLoginVerificationMiddleware() *LoginVerificationMiddleware {
	return &LoginVerificationMiddleware{}
}

func (m *LoginVerificationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		if r.Header.Get("Token") == "xin000000" {
			next(w, r)
			return
		}
		w.Write([]byte("权限不足"))
		return
	}
}
