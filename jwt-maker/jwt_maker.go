/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package jwt_maker

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/url"
)

func MakeJwtWithoutPayload() string {
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

func MakeJwtWithPayload(payload url.Values) string {

	claims := make(jwt.MapClaims)
	claims["access_key"] = accessKey
	claims["nonce"] = uuid.NewV4().String()
	claims["query"] = payload.Encode()
	qh := sha512.New()
	qh.Reset()
	qh.Write([]byte(payload.Encode()))
	claims["query_hash"] = hex.EncodeToString(qh.Sum(nil))
	claims["query_hash_alg"] = "SHA512"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	//log.Println(signedToken)
	if err != nil {
		log.Println("SignedString Error")
		log.Fatal(err)
		return ""
	}
	return signedToken
}
