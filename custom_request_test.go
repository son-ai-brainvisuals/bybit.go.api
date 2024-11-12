package bybit_connector

import (
	"context"
	"net/url"
	"testing"

	"log"
)

func TestSendCustomSignRequestWalletBalance(t *testing.T) {

	API_KEY := "63dXtR1GFmrD6ywPGc"
	API_SECRET := "vzu7tehZytvqCJxve9I8ma5pdYBGE8U9FcvF"
	opts := []ClientOption{}
	opts = append(opts, WithBaseURL(TESTNET))
	c := NewBybitHttpClient(API_KEY, API_SECRET, opts...)

	query := &url.Values{}
	query.Set("accountType", "UNIFIED")
	res, err := c.SendCustomSignRequest(context.TODO(), "GET", "/v5/account/wallet-balance", query, nil)
	if err != nil {
		t.Fatal(err)
	}
	resBytes, _ := json.Marshal(res)
	log.Println("res", string(resBytes))
}

const apiKey = "M86Bz93UwsidYn70cQ"
const apiSecret = "9FrlRwLWTcOsES35DEFK0nfUzsdiCEH7kguU"

func TestSendCustomSignRequestCopyTrading(t *testing.T) {

	opts := []ClientOption{}
	opts = append(opts, WithBaseURL(TESTNET))
	c := NewBybitHttpClient(apiKey, apiSecret, opts...)

	query := &url.Values{}
	query.Set("accountType", "UNIFIED")
	res, err := c.SendCustomSignRequest(context.TODO(), "GET", "/v5/account/wallet-balance", query, nil)
	if err != nil {
		t.Fatal(err)
	}
	resBytes, _ := json.Marshal(res)
	log.Println("res", string(resBytes))
}
