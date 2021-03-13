/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package main

import (
	"encoding/json"
	"github.com/shieldnet/gobit/strategy"
	"io/ioutil"
	"log"
	"os"
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

type CandleSet struct {
	Bc int `json:"bc"`
	Sc int `json:"sc"`
}

func main() {

	jsonFile, err := os.Open("result.json")
	if err != nil{
		log.Panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cinfo map[string]CandleSet
	json.Unmarshal(byteValue, &cinfo)

	coinList := []string{Tfuel, Ankr, Chz, Pxl, Kava, Solve, Mlk, Obsr}
	var strategies []*strategy.Strategy
	for _, coin := range coinList {
		s := &strategy.Strategy{
			Market:        coin,
			BuyCandleNum:  cinfo[coin].Bc,
			SellCandleNum: cinfo[coin].Sc,
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
