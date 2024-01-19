package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateManagementToken(appAccessKey, appSecret string) (string, error) {
	if appAccessKey == "" || appSecret == "" {
		return "", errors.New("missing app key or secret")
	}
	mySigningKey := []byte(appSecret)
	expiresIn := uint32(24 * 3600)
	now := uint32(time.Now().UTC().Unix())
	exp := now + expiresIn
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": appAccessKey,
		"type":       "management",
		"version":    2,
		"jti":        uuid.New().String(),
		"iat":        now,
		"exp":        exp,
		"nbf":        now,
	})

	// Sign and get the complete encoded token as a string using the secret
	signedToken, _ := token.SignedString(mySigningKey)
	return signedToken, nil
}
