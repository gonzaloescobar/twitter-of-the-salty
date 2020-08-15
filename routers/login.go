package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gonzaloescobar/twitter-of-the-salty/db"
	"github.com/gonzaloescobar/twitter-of-the-salty/jwt"
	"github.com/gonzaloescobar/twitter-of-the-salty/models"
)

/*Login to app*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Invalid user or password "+err.Error(), 400)
		return
	}

	if len(u.Mail) == 0 {
		http.Error(w, "The mail is mandatory ", 400)
		return
	}
	document, exist := db.Login(u.Mail, u.Password)
	if exist == false {
		http.Error(w, "Invalid user or password ", 400)
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error while generate Token "+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "applidaction/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
