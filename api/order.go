/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package api

import (
	"github.com/shieldnet/gobit/jwtmaker"
	"log"
	"net/url"
)

const (
	OrderURL = "https://api.upbit.com/v1/orders"
)

func BuyOrderByMarketPrice(market, totalPrice string, key jwtmaker.Keys) string {
	ord_type := "price"
	side := "bid"
	data := makeUrlValues(market, side, "", totalPrice, ord_type)
	header := map[string]string{
		"Authorization": "Bearer " + jwtmaker.MakeJwtWithPayload(key, data),
	}
	resp, err := HttpPOST(OrderURL, header, data)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return string(resp)
}

func SellOrderByMarketPrice(market, volume string, key jwtmaker.Keys) string {
	ord_type := "market"
	side := "ask"
	data := makeUrlValues(market, side, volume, "", ord_type)
	header := map[string]string{
		"Authorization": "Bearer " + jwtmaker.MakeJwtWithPayload(key, data),
	}
	resp, err := HttpPOST(OrderURL, header, data)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return string(resp)
}

func makeUrlValues(market, side, volume, price, ord_type string) url.Values {
	data := url.Values{}
	if market != "" {
		data["market"] = []string{market}
	}
	if market != "" {
		data["side"] = []string{side}
	}
	if volume != "" {
		data["volume"] = []string{volume}
	}
	if price != "" {
		data["price"] = []string{price}
	}
	if ord_type != "" {
		data["ord_type"] = []string{ord_type}
	}
	return data
}
