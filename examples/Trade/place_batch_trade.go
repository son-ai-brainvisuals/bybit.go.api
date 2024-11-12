package main

import (
	"context"
	"fmt"
	bybit "github.com/son-ai-brainvisuals/bybit.go.api"
)

func main() {
	PlaceBatchTrade()
}

func PlaceBatchTrade() {
	client := bybit.NewBybitHttpClient("8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "option",
		"request": []map[string]interface{}{
			{
				"category":    "option",
				"symbol":      "BTC-10FEB23-24000-C",
				"orderType":   "Limit",
				"side":        "Buy",
				"qty":         "0.1",
				"price":       "5",
				"orderIv":     "0.1",
				"timeInForce": "GTC",
				"orderLinkId": "9b381bb1-401",
				"mmp":         false,
				"reduceOnly":  false,
			},
			{
				"category":    "option",
				"symbol":      "BTC-10FEB23-24000-C",
				"orderType":   "Limit",
				"side":        "Buy",
				"qty":         "0.1",
				"price":       "5",
				"orderIv":     "0.1",
				"timeInForce": "GTC",
				"orderLinkId": "82ee86dd-001",
				"mmp":         false,
				"reduceOnly":  false,
			},
		},
	}
	orderResult, err := client.NewTradeService(params).PlaceBatchOrder(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
