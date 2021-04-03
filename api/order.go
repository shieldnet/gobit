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
	resp := Order(market, "bid", "", totalPrice, "price", "", key)
	return resp
}

func SellOrderByMarketPrice(market, volume string, key jwtmaker.Keys) string {
	resp := Order(market, "ask", volume, "", "market", "", key)
	return resp
}

func Order(market, side, volume, price, ordType, identifier string, key jwtmaker.Keys) string {
	data := makeUrlValues(market, side, volume, price, ordType, identifier)
	header := map[string]string{
		"Authorization": "Bearer " + jwtmaker.MakeJwtWithPayload(key, data),
	}
	resp, err := HttpPOST(OrderURL, header, data)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return string(resp)
}

func Buy(market, price, volume string, key jwtmaker.Keys) string {
	return Order(market, "ask", volume, price, "limit", "", key)
}

func Sell(market, price, volume string, key jwtmaker.Keys) string {
	return Order(market, "bid", volume, price, "limit", "", key)
}

func Cancle(uuid, identifier string, key jwtmaker.Keys) string {
	data := url.Values{}
	if uuid != "" {
		data["uuid"] = []string{uuid}
	}
	if identifier != "" {
		data["identifier"] = []string{identifier}
	}
	header := map[string]string{
		"Authorization": "Bearer " + jwtmaker.MakeJwtWithPayload(key, data),
	}
	resp, err := HttpDelete(OrderURL, header, data)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return string(resp)
}

func GetOrderList(market string, states []string, key jwtmaker.Keys) string {
	data := url.Values{}
	if market != "" {
		data["market"] = []string{market}
	}
	if len(states) > 0 {
		data["market"] = states
	}
	header := map[string]string{
		"Authorization": "Bearer " + jwtmaker.MakeJwtWithPayload(key, data),
	}
	resp, err := HttpGet(OrderURL, header, data)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return string(resp)
}

func GetOrderInfo(uuid, identifier string, key jwtmaker.Keys) string {
	data := url.Values{}
	if uuid != "" {
		data["uuid"] = []string{uuid}
	}
	if identifier != "" {
		data["identifier"] = []string{identifier}
	}
	header := map[string]string{
		"Authorization": "Bearer " + jwtmaker.MakeJwtWithPayload(key, data),
	}
	resp, err := HttpGet(OrderURL, header, data)
	if err != nil {
		log.Fatalln(string(resp), err)
	}
	return string(resp)
}

func makeUrlValues(market, side, volume, price, ord_type, identifier string) url.Values {
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
	if identifier != "" {
		data["identifier"] = []string{identifier}
	}
	return data
}
