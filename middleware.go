package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func jwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("Lindu"), nil
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		_, ok := token.Claims.(*myClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token invalid"))
			return
		}
		next(w, r)
	}
}
