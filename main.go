package main

import "github.com/shieldnet/gobit/strategy"

const (
	Tfuel = "KRW-TFUEL"
	MLK = "KRW-PCI"
)

func main() {
	strategy.Init(MLK)
}