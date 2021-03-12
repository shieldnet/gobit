package strategy

import (
	"github.com/shieldnet/gobit/api"
	"log"
	"strconv"
	"strings"
	"time"
)

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

func Init(market string) {
	log.Println("[Init] 전략에 진입합니다." + time.Now().String())
	time.Sleep(500 * time.Millisecond)
	BuyCheck(market)
}

func BuyCheck(market string) {
	for true {
		candles := api.GetMinuteCandle(3, 15, market)
		nowTradeValue := candles[0].TradePrice
		if isHighestTradeValue(nowTradeValue, candles) {
			log.Println("[BuyCheck] 28개중 최고가이므로 구매" + time.Now().String())
			Buy(market)
			break
		}
		time.Sleep(500 * time.Millisecond)
		log.Println("[BuyCheck] 조건을 만족하지 않았으므로, 구매하지 않습니다." + time.Now().String())
	}
}

func Buy(market string) {
	log.Println("[Buy] 구매를 시작합니다." + time.Now().String())
	api.BuyOrderByMarketPrice(market, "200000")
	time.Sleep(500 * time.Millisecond)
	SellCheck(market)
}

func SellCheck(market string) {
	log.Println("[SellCheck] 주식을 팔지 말지 생각해봅니다." + time.Now().String())
	for true {
		candles := api.GetMinuteCandle(3, 15, market)
		accounts := api.GetAccountInfo()
		nowTradeValue := candles[0].TradePrice
		avgBuyPrice := 0.0
		balance := ""
		for _, account := range accounts {
			log.Println(account.Currency, market, account.Balance)
			if account.Currency == strings.ReplaceAll(market,"KRW-", "") {
				avgBuyPrice, _ = strconv.ParseFloat(account.AvgBuyPrice, 32)
				balance = account.Balance
				break
			}
		}

		time.Sleep(500 * time.Millisecond)

		if isQuitPrice(float64(nowTradeValue), avgBuyPrice, 2.3) {
			log.Println("[SellCheck] 손절률 이상 가격이 떨어졌으므로, 손절합니다." + time.Now().String())
			Sell(market, balance)
			break
		}
		if isLowestTradeValue(nowTradeValue, candles) {
		//if true {
			log.Println("[SellCheck] 28개중 최저가이므로 판매합니다." + time.Now().String())
			Sell(market, balance)
			break
		}
		time.Sleep(500 * time.Millisecond)
		log.Println("[SellCheck] 조건을 만족하지 않았으므로, 판매하지 않습니다." + time.Now().String())
	}
}

func Sell(market string, balance string) {
	//println(market, balance)
	println(api.SellOrderByMarketPrice(market, balance))
}
