package jwt_maker

import (
	"crypto/sha512"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"log"
)

// payload 있으면 args = `{"key":"value"}`, 없으면 nil
func MakeJWT(args interface{}) string {
	switch args.(type) {
	case string:
		str := fmt.Sprintf("%v", args)
		return makeJwtWithPayload(str)
	case nil:
		return makeJwtWithoutPayload()
	default:
		return ""
	}
	return ""
}

func makeJwtWithoutPayload() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["access_key"] = accessKey
	claims["nonce"] = uuid.NewV4().String()
	token.Claims = claims

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("SignedString Error")
		log.Fatal(err)
		return ""
	}
	return signedToken
}

func makeJwtWithPayload(payload string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["access_key"] = accessKey
	claims["query"] = payload
	claims["query_hash"] = sha512.New().Sum([]byte(payload))
	claims["query_hash_algorithm"] = "SHA512"
	token.Claims = claims

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("SignedString Error")
		log.Fatal(err)
		return ""
	}
	return signedToken
}
