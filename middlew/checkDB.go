package middlew

import (
	"net/http"

	"github.com/gonzaloescobar/twitter-of-the-salty/db"
)

/*CheckDB status db*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Lost connection DB, 500", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
