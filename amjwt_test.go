package amjwt_test

import (
	"github.com/golang-jwt/jwt"
	"github.com/yukitsune/amjwt"
	"io/ioutil"
	"testing"
	"time"
)

const testKey = "MyTestKey"
const testTeamId = "MyTestTeamId"
const defaultTestExpiry = 30 * 6

func TestJwtContainsKeyId(t *testing.T) {

	privateKey, publicKey, err := readTestKeys()
	if err != nil {
		t.Error(err)
		return
	}

	token, err := amjwt.CreateJwt(testKey, testTeamId, defaultTestExpiry, privateKey)
	if err != nil {
		t.Error(err)
		return
	}

	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if tok.Header["kid"] != testKey {
		t.Fail()
		return
	}
}

func TestJwtContainsTeamId(t *testing.T) {

	privateKey, publicKey, err := readTestKeys()
	if err != nil {
		t.Error(err)
		return
	}

	token, err := amjwt.CreateJwt(testKey, testTeamId, defaultTestExpiry, privateKey)
	if err != nil {
		t.Error(err)
		return
	}

	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	claimsMap := tok.Claims.(jwt.MapClaims)

	if claimsMap["iss"] != testTeamId {
		t.Fail()
		return
	}
}

func TestJwtExpiryIsSet(t *testing.T) {

	privateKey, publicKey, err := readTestKeys()
	if err != nil {
		t.Error(err)
		return
	}

	// 5 day expiry
	expiry := 5

	token, err := amjwt.CreateJwt(testKey, testTeamId, expiry, privateKey)
	if err != nil {
		t.Error(err)
		return
	}

	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	currentTime := time.Now().Unix()

	claimsMap := tok.Claims.(jwt.MapClaims)
	expiryTime := claimsMap["exp"].(float64)

	diff := int64(expiryTime) - currentTime
	expectedDiff := int64(expiry * 24 * 60 * 60)

	// It's really unlikely that we need the buffers here, but i'm paranoid anyway...
	if diff > expectedDiff + 1 || diff < expectedDiff - 1 {
		t.Error(err)
		return
	}
}

func readTestKeys() (privateKey []byte, publicKey []byte, err error) {
	privateKey, err = ioutil.ReadFile("test/private_key.pem")
	if err != nil {
		return privateKey, publicKey, err
	}

	publicKey, err = ioutil.ReadFile("test/private_key.pem")
	return privateKey, publicKey, err
}
