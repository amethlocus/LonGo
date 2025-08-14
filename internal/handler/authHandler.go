package handler

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: Make this a global variable
var secretkey = []byte("secret-key")

func CreateJWT(username string) (string, error) {
	// claims creation
	claimsMap := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 200).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMap)

	// adds signature to token
	tokenstring, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}

	VerifyToken(tokenstring)
	return tokenstring, nil
}

func VerifyToken(tokenstring string) error {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	fmt.Printf("Parsed and verified %v - %v", token.Valid, token.Raw)
	return nil
}
