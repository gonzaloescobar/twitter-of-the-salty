package db

import (
	"github.com/gonzaloescobar/twitter-of-the-salty/models"
	"golang.org/x/crypto/bcrypt"
)

/*Login check db*/
func Login(mail string, password string) (models.User, bool) {
	user, exist, _ := CheckUserAlreadyExists(mail)
	if exist == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
