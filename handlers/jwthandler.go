package handlers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// func verifyToken(c *fiber.Ctx) {

// }

// func CreateToken(message string) {

// 	loginstring := message

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"msg": loginstring,
// 		"nbf": "LocalUser" + "earld@winjit.com",
// 		"exp":time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
// 	})

// 	var secret = []byte("testing")

// 	tokenString, err := token.SignedString(secret)
// 	if err != nil {
// 	}

// }

// func VerifyToken(message string) {

// }
// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken(usernameincoming string, roleid int) (tokenout string) {
	// Set secret key from .env file.
	secret := "secret"

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi("10")

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	//claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["username"] = usernameincoming
	claims["role"] = roleid

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		//return "", err
	}
	tokenstring := t

	return tokenstring
	//return t, nil

}

//this is my logic to verify tokens
func VerifyToken(tokenString string) (tokenFlag bool) {
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		//
	}

	if token.Valid {
		tokenFlag = true
		return tokenFlag
	} else {
		tokenFlag = false
		return tokenFlag
	}
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte("secret"), nil
}
