package main

import (
	"github.com/shieldnet/gobit/api"
	jwtmaker "github.com/shieldnet/gobit/jwt-maker"
)

func main() {
	//token := jwtmaker.MakeJWT(`{"key":"value"}`)
	token2 := jwtmaker.MakeJWT(nil)
	//println(token)

	println(token2)
	cd := api.GetMinuteCandle(5, 200, "KRW-XRP")
	println(cd[0].OpeningPrice)
}