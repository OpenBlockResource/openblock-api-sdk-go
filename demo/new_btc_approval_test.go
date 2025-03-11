package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestNewBtcApproval(t *testing.T) {
	client := apisdk.NewClient("361bdf3a1e0640979a3e2240c3361609", "U3IXmFgR848q5XA1vQVgdvW1Z69UvQXD", 10*time.Second)

	utxo := apisdk.Utxo{
		Script: "76a914f6b3f7f1f1b1b1b1b1b1b1b1b1b1b1b1b1b1b1b188ac",
		Amount: "0.001",
		Hash:   "0x7620E9bbb6591b9F53b6b5002B716B8D23ca3950",
		Index:  "0",
	}

	resp, err := client.CompanyWallet.NewApproval(&apisdk.ParamNewApproval{
		Action: "TRANSACTION",
		TXInfo: apisdk.TXInfo{
			To:              "bc1qmax35qkxgke2aq2xugaxjl96q8f0k7l3ndhz9g",
			From:            "bc1qmax35qkxgke2aq2xugaxjl96q8f0k7l3ndhz9g",
			Value:           "0.001",
			GasPrice:        "16500",
			Chain:           "BTC",
			TransactionType: "native",
			Utxo:            []*apisdk.Utxo{&utxo},
		},
	})
	fmt.Println(resp, err)
}
