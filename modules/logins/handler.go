package logins

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) Login(w http.ResponseWriter, r *http.Request) {

	var user User
	json.NewDecoder(r.Body).Decode(&user)
	token, err := handler.Usecase.Login(user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Login Failed"))
		} else {
			w.Write([]byte(err.Error()))
		}
		return
	}
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println("token :", token)
	w.Write([]byte(token))
}
