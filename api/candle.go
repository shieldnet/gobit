/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	ApiAddr = "https://api.upbit.com/v1/"
)

const (
	Candles = "candles/"
)

const (
	ByMinutes = "minutes/"
	ByDays = "days"
	ByWeeks = "weeks"
	ByMonths = "months"
)

type Candle struct {
	Market               string  `json:"market"`
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`
	CandleDataTimeKST    string  `json:"candle_date_time_kst"`
	OpeningPrice         float32 `json:"opening_price"`
	HighPrice            float32 `json:"high_price"`
	LowPrice             float32 `json:"low_price"`
	TradePrice           float32 `json:"trade_price"`
	Timestamp            uint64  `json:"timestamp"`
	CandleAccTradePrice  float32 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float32 `json:"candle_acc_trade_volume"`
	Unit                 int     `json:"unit"`
}

type CandleList []Candle

func GetMinuteCandle(unit, count int, market string) CandleList {
	req, err := http.NewRequest("GET", ApiAddr+Candles+ByMinutes+strconv.Itoa(unit), nil)
	if err != nil {
		log.Panic(err)
	}
	q := req.URL.Query()
	q.Add("market", market)
	q.Add("count", strconv.Itoa(count))
	req.URL.RawQuery = q.Encode()

	resp, _ := HttpGet(req.URL.String(), map[string]string{}, url.Values{})
	//println(string(resp))
	candles := CandleList{}
	err = json.Unmarshal(resp, &candles)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return candles
}

func GetAllMinuteCandlesOfMarket(unit, count int) map[string]CandleList {
	ret := map[string]CandleList{}
	markets := GetAllMarketCodes()
	for _, market := range markets {
		marketCode := market.Market // 종목 코드
		ret[marketCode] = GetMinuteCandle(unit, count, marketCode)
		time.Sleep(110 * time.Millisecond) // Quotation은 IP당 1초에 10개니까 API 개수 제한 걸자.
	}
	return ret
}

func GetDayCandle(count int, market string) CandleList {
	req, err := http.NewRequest("GET", ApiAddr+Candles+ByDays, nil)
	if err != nil {
		log.Panic(err)
	}
	q := req.URL.Query()
	q.Add("market", market)
	q.Add("count", strconv.Itoa(count))
	req.URL.RawQuery = q.Encode()

	resp, _ := HttpGet(req.URL.String(), map[string]string{}, url.Values{})
	//println(string(resp))
	candles := CandleList{}
	err = json.Unmarshal(resp, &candles)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return candles
}

func GetWeekCandle(count int, market string) CandleList {
	req, err := http.NewRequest("GET", ApiAddr+Candles+ByDays, nil)
	if err != nil {
		log.Panic(err)
	}
	q := req.URL.Query()
	q.Add("market", market)
	q.Add("count", strconv.Itoa(count))
	req.URL.RawQuery = q.Encode()

	resp, _ := HttpGet(req.URL.String(), map[string]string{}, url.Values{})
	//println(string(resp))
	candles := CandleList{}
	err = json.Unmarshal(resp, &candles)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return candles
}

func GetMonthCandle(count int, market string) CandleList {
	req, err := http.NewRequest("GET", ApiAddr+Candles+ByMonths, nil)
	if err != nil {
		log.Panic(err)
	}
	q := req.URL.Query()
	q.Add("market", market)
	q.Add("count", strconv.Itoa(count))
	req.URL.RawQuery = q.Encode()

	resp, _ := HttpGet(req.URL.String(), map[string]string{}, url.Values{})
	//println(string(resp))
	candles := CandleList{}
	err = json.Unmarshal(resp, &candles)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return candles
}