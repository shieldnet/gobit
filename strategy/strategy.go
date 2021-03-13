package strategy

import (
	"github.com/shieldnet/gobit/api"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Trader interface {
	Execute()
	Init(string)
	BuyCheck(string)
	Buy(string)
	SellCheck(string)
	Sell(string)
}

type Strategy struct {
	Market        string
	BuyCandleNum  int
	SellCandleNum int
	NextState     string
	CandleUnit    int
	QuitRate      float64
	Balance       string
	TotalPrice    string
}

func (s *Strategy) Execute(wait *sync.WaitGroup) {
	defer wait.Done()

	f := map[string]func(){
		"Init":      s.Init,
		"BuyCheck":  s.BuyCheck,
		"Buy":       s.Buy,
		"SellCheck": s.SellCheck,
		"Sell":      s.Sell,
	}
	f[s.NextState]()
	return
}

func (s *Strategy) Init() {
	log.Println("[Init] 전략에 진입합니다." + s.Market)
	time.Sleep(SleepInterval * time.Millisecond)
	s.NextState = "BuyCheck"
	return
}

func (s *Strategy) BuyCheck() {
	candles := api.GetMinuteCandle(3, s.BuyCandleNum, s.Market)
	nowTradeValue := candles[0].TradePrice
	if isHighestTradeValue(nowTradeValue, candles) {
		log.Println("[BuyCheck] 15개중 최고가이므로 구매" + s.Market)
		s.NextState = "Buy"
		return
	}
	time.Sleep(SleepInterval * time.Millisecond)
	log.Println("[BuyCheck] 조건을 만족하지 않았으므로, 구매하지 않습니다." + s.Market)

	return
}

func (s *Strategy) Buy() {
	log.Println("[Buy] 구매를 시작합니다." + s.Market)
	api.BuyOrderByMarketPrice(s.Market, s.TotalPrice)
	time.Sleep(SleepInterval * time.Millisecond)
	s.NextState = "SellCheck"
	return
}

func (s *Strategy) SellCheck() {
	log.Println("[SellCheck] 주식을 팔지 말지 생각해봅니다." + s.Market)

	candles := api.GetMinuteCandle(s.CandleUnit, s.SellCandleNum, s.Market)
	accounts := api.GetAccountInfo()
	nowTradeValue := candles[0].TradePrice
	avgBuyPrice := 0.0
	for _, account := range accounts {
		//log.Println(account.Currency, market, account.Balance)
		if account.Currency == strings.ReplaceAll(s.Market, "KRW-", "") {
			avgBuyPrice, _ = strconv.ParseFloat(account.AvgBuyPrice, 32)
			s.Balance = account.Balance
			break
		}
	}

	if isQuitPrice(float64(nowTradeValue), avgBuyPrice, s.QuitRate) {
		log.Println("[SellCheck] 손절률 이상 가격이 떨어졌으므로, 손절합니다." + s.Market)
		s.NextState = "Sell"
		return
	}
	if isLowestTradeValue(nowTradeValue, candles) {
		log.Println("[SellCheck] 7개중 최저가이므로 판매합니다." + s.Market)
		s.NextState = "Sell"
		return
	}
	time.Sleep(SleepInterval * time.Millisecond)
	log.Println("[SellCheck] 조건을 만족하지 않았으므로, 판매하지 않습니다." + s.Market)

	return
}

func (s *Strategy) Sell() {
	api.SellOrderByMarketPrice(s.Market, s.Balance)
	s.NextState = "Init"
	return
}
