package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func VerifyJWT(encodedToken string, secretKey string) (*jwt.Token, *jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token alg: %s", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return token, &jwt.RegisteredClaims{}, errors.Wrap(err, "parse with claims")
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		claims = &jwt.RegisteredClaims{}
	}

	return token, claims, nil
}
