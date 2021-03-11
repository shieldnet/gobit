package api

import (
	jwt_maker "github.com/shieldnet/gobit/jwt-maker"
	"net/url"
)

const (
	OrderURL = "https://api.upbit.com/v1/orders"
)

func BuyOrderByMarketPrice(market, totalPrice string) string {
	ord_type := "price"
	side := "bid"
	data := makeUrlValues(market, side, "", totalPrice, ord_type)
	header := map[string]string{
		"Authorization": "Bearer " + jwt_maker.MakeJwtWithPayload(data),
	}
	resp, _ := HttpPOST(OrderURL, header, data)
	return string(resp)
}

func SellOrderByMarketPrice(market, volume string) string {
	ord_type := "market"
	side := "ask"
	data := makeUrlValues(market, side, volume, "", ord_type)
	header := map[string]string{
		"Authorization": "Bearer " + jwt_maker.MakeJwtWithPayload(data),
	}
	resp, _ := HttpPOST(OrderURL, header, data)
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