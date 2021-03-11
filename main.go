package main

import "github.com/shieldnet/gobit/strategy"

const (
	Tfuel = "KRW-TFUEL"
	MLK = "KRW-ANKR"
)

func main() {
	strategy.Init(MLK)

}