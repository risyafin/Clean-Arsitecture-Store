package logins

import "github.com/golang-jwt/jwt"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
