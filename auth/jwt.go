package Auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type JwtAuth struct {
}

func NewJwtAuth() *JwtAuth {
	return &JwtAuth{}
}

func (jwtAuth *JwtAuth) GetToken(email string) (string, error) {

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	if err != nil {
		return "", errors.New("Failed to generate token")
	}

	return tokenString, err
}

func (jwtAuth *JwtAuth) ValidateToken(tokenString string) (bool, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used to sign the token
		return []byte(viper.GetString("secret")), nil
	})
	if err != nil {
		return false, nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims, nil
	} else {
		return false, nil, nil
	}
}
