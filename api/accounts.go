/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package api

import (
	"encoding/json"
	"github.com/shieldnet/gobit/jwtmaker"
	"log"
	"net/http"
	"net/url"
)

type Account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

type AccountList []Account

func GetAccountInfo(key jwtmaker.Keys) AccountList {
	req, err := http.NewRequest("GET", "https://api.upbit.com/v1/accounts", nil)
	if err != nil {
		log.Panic(err)
	}
	//println(req.URL.String())

	resp, _ := HttpGet(req.URL.String(), map[string]string{"Authorization": "Bearer " + jwtmaker.MakeJwtWithoutPayload(key)}, url.Values{})
	//println(string(resp))
	accounts := AccountList{}
	err = json.Unmarshal(resp, &accounts)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return accounts
}
