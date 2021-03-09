package main

import (
	jwtmaker "github.com/shieldnet/gobit/jwt-maker"
)

func main() {
	token := jwtmaker.MakeJWT(`{"key":"value"}`)
	token2 := jwtmaker.MakeJWT(nil)
	println(token)
	println(token2)
}