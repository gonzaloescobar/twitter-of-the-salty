package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gonzaloescobar/twitter-of-the-salty/db"
	"github.com/gonzaloescobar/twitter-of-the-salty/models"
)

/*Mail value used for all endponts*/
var Mail string

/*IDUser value used for all endponts*/
var IDUser string

/*ProcessToken : process token and get values*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	key := []byte("LosGrandesNoDescienden")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, exist, _ := db.CheckUserAlreadyExists(claims.Mail)
		if exist == true {
			Mail = claims.Mail
			IDUser = claims.ID.Hex()
		}
		return claims, exist, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}
	return claims, false, string(""), err
}
