package routers

import (
	"encoding/json"
	"net/http"
	"github.com/gonzaloescobar/twitter-of-the-salty/db"
	"github.com/gonzaloescobar/twitter-of-the-salty/models"
)

func User(w http.ResponseWriter, r *http.Request){

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in received data"+err.Error(), 400)
		return
	}

	if len(t.Mail) == 0 {
		http.Error(w, "Mail mandatory", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "The password must be at least 6 characters", 400)
		return
	}

	_, exist, _ := db.CheckUserAlreadyExists(t.Mail)
	if exist == true {
		http.Error(w, "User Already exists with that mail", 400)
		return
	}

	_, status, err := db.InserUser(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to enrollment user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w,"User insert failed", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}