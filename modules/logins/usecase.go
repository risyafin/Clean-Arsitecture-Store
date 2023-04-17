package logins

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) Login(user User) (string, error) {
	err := usecase.Repo.Login(user)
	claims := MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("Lindu"))
	

	return signedToken,err
}
