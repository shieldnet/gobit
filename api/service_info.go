package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Market struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
}

type MarketList []Market

func GetAllMarketCodes() MarketList {
	req, err := http.NewRequest("GET", ApiAddr+"market/all", nil)
	if err != nil {
		log.Panic(err)
	}
	q := req.URL.Query()
	q.Add("isDetails", "false")
	req.URL.RawQuery = q.Encode()

	resp, _ := HttpGet(req.URL.String(), map[string]string{})

	markets := MarketList{}
	err = json.Unmarshal(resp, &markets)
	if err != nil {
		log.Fatalln(string(resp),err)
	}
	return markets
}
