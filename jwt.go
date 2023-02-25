package core

import (	
    "time"
	"errors"
    "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type JwtAuth struct{

}

func NewJwtAuth(){
	return &JwtAuth{}
}

func (*jwtAuth JwtAuth) GetToken(email string) (string, error){

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	if err != nil {
		return "", errors.New("Failed to generate token")
	}

	return tokenString, err
}