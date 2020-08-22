package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gonzaloescobar/twitter-of-the-salty/db"
)

/*ViewProfile : Allows you to extract the values ​​from the profile*/
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to find a profile "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
