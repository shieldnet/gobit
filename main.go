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
)

func main() {
	coinList := []string{Tfuel, Ankr, Chz, Meta, Pxl, Kava, Med}
	var strategies []*strategy.Strategy
	for _, coin := range coinList {
		s := &strategy.Strategy{
			Market:        coin,
			BuyCandleNum:  15,
			SellCandleNum: 8,
			QuitRate:      2.0,
			CandleUnit:    3,
			NextState:     "Init",
			Balance:       "0",
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
