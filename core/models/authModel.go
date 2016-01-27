package models

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"github.com/remony/Equipment-Rental-API/core/database"
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
	"regexp"
)

type Error_response struct {
	Message string `json:"message"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func authLogin(password string, digest string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}

func isValidEntry(username string) bool {
	if len(username) < 240 {
		if ok, _ := regexp.MatchString("^[A-Za-z0-9]+$", username); ok {
			return true
		} else {
			return false
		}
	}
	return false
}

func PerformLogin(api router.API, username string, password string) database.Auth {
	var digest = database.GetDigest(api, username)
	var login database.Auth
	if isValidEntry(username) && isValidEntry(password) {
		if (authLogin(password, digest)) {
			login = database.LoginUser(api, strings.ToLower(username))
			return login;

		} else {
			return database.Auth{
				Success: false,
			}
		}
	} else {
		return database.Auth{
			Success: false,
		}
	}

}

func PerformRegister(api router.API, data Register) bool {
	if !database.CheckIfUserExists(api, data.Username) {
		if (secureEntry(data.Password) && secureEntry(data.Username) && secureEntry(data.Email)) {

			hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
			if err != nil {
				log.Fatal(err)
			}

			if database.RegisterUser(api, data.Username, hash, data.Email) {
				return true;
			} else {
				return false;
			}
		}

	}
	return false
}


// secureEntry guards the api from entering data which is not acceptable
func secureEntry(password string) bool {
	// if the value if more or equal to 6 || if the value does not contain any spaces
	if len(password) >= 6 && len(strings.Split(password, " ")) == 1 {
		// Return true that it is acceptable
		return true;
	}
	// Return false that it should not be accepted
	return false
}

func PerformRemoveUser(api router.API, data Register) bool {
	if (database.CheckIfUserExists(api, data.Username)) {
		database.RemoveUser(api, data.Username)
		return true
	}
	return false
}

