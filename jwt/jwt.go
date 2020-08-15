package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gonzaloescobar/twitter-of-the-salty/models"
)

/*GenerateJWT Generate token */
func GenerateJWT(u models.User) (string, error) {
	key := []byte("LosGrandesNoDescienden")
	payload := jwt.MapClaims{
		"email":      u.Mail,
		"name":       u.FirstName,
		"lastname":   u.LastName,
		"birth_date": u.BirthDate,
		"biography":  u.Biography,
		"location":   u.Location,
		"website":    u.Website,
		"_id":        u.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return tokenString, err
	}
	return tokenString, err
}
