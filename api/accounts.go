/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package api

import (
	"encoding/json"
	jwt_maker "github.com/shieldnet/gobit/jwt-maker"
	"log"
	"net/http"
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

func GetAccountInfo() AccountList {
	req, err := http.NewRequest("GET", "https://api.upbit.com/v1/accounts", nil)
	if err != nil {
		log.Panic(err)
	}
	//println(req.URL.String())

	resp, _ := HttpGet(req.URL.String(), map[string]string{"Authorization": "Bearer " + jwt_maker.MakeJwtWithoutPayload(jwt_maker.MainKey)})
	//println(string(resp))
	accounts := AccountList{}
	err = json.Unmarshal(resp, &accounts)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return accounts
}
