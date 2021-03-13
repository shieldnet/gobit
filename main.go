package main

import (
	"github.com/shieldnet/gobit/strategy"
	"sync"
)

const (
	Tfuel = "KRW-TFUEL"
	Ankr  = "KRW-ANKR"
	Chz   = "KRW-CHZ"
	Meta  = "KRW-META"
	Pxl   = "KRW-PXL"
	Kava  = "KRW-KAVA"
	Med   = "KRW-MED"
	Qkc   = "KRW-QKC"
	Mlk   = "KRW-MLK"
	Btc   = "KRW-BTC"
	Solve = "KRW-SOLVE"
	Obsr  = "KRW-OBSR"
)

func main() {
	coinList := []string{Tfuel, Ankr, Chz, Pxl, Kava, Solve, Mlk, Obsr}
	var strategies []*strategy.Strategy
	for _, coin := range coinList {
		s := &strategy.Strategy{
			Market:        coin,
			BuyCandleNum:  20,
			SellCandleNum: 20,
			QuitRate:      2.0,
			CandleUnit:    5,
			NextState:     "Init",
			Balance:       "0",
			TotalPrice:    "50000",
		}
		strategies = append(strategies, s)
	}

	for true {
		println("-------------------상황판--------------------")
		wait := new(sync.WaitGroup)
		wait.Add(len(strategies))
		for _, s := range strategies {
			go s.Execute(wait)
		}
		wait.Wait()
		println("---------------------------------------------")
	}
}
