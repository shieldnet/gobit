/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package strategy

import (
	"github.com/shieldnet/gobit/api"
)

const SleepInterval = 1500

func isHighestTradeValue(nowTradeValue float32, candles api.CandleList) bool {
	for _, candle := range candles[1:] {
		if candle.TradePrice > nowTradeValue {
			return false
		}
	}
	return true
}

func isLowestTradeValue(nowTradeValue float32, candles api.CandleList) bool {
	for _, candle := range candles[1:] {
		if candle.TradePrice < nowTradeValue {
			return false
		}
	}
	return true
}

func isQuitPrice(nowTradeValue, avgBuyPrice, quitPercent float64) bool {
	if nowTradeValue <= avgBuyPrice*((100.0-quitPercent)/100.0) {
		return true
	}
	return false
}