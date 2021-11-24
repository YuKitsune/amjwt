package amjwt

import (
	"crypto/ecdsa"
	"github.com/golang-jwt/jwt"
	"time"
)

// CreateJwt will create the JWT and return it as a string
func CreateJwt(keyId string, teamId string, expiryDays int, privateKeyBytes []byte) (string, error) {

	issueTime := time.Now()
	expireTime := issueTime.AddDate(0, 0, expiryDays)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": teamId,
		"iat": issueTime.Unix(),
		"exp": expireTime.Unix(),
	})

	token.Header["kid"] = keyId

	var privateKey *ecdsa.PrivateKey
	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
