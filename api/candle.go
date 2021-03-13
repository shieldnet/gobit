package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const (
	ApiAddr = "https://api.upbit.com/v1/"
)

const (
	Candles = "candles/"
)

const (
	ByMinutes = "minutes/"
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
		panic(err)
	}
	q := req.URL.Query()
	q.Add("market", market)
	q.Add("count", strconv.Itoa(count))
	req.URL.RawQuery = q.Encode()

	resp, _ := HttpGet(req.URL.String(), map[string]string{})
	//println(string(resp))
	candles := CandleList{}
	err = json.Unmarshal(resp, &candles)
	if err != nil {
		log.Fatalln(string(resp),err)
	}
	return candles
}
