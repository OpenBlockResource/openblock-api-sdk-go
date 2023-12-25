package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestNewApproval(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.NewApproval(&apisdk.ParamNewApproval{
		Action: "TRANSACTION",
		TXInfo: apisdk.TXInfo{
			To:                   "0x7620E9bbb6591b9F53b6b5002B716B8D23ca3950",
			TokenAddress:         "0xa71edc38d189767582c38a3145b5873052c3e47a",
			From:                 "0x5adE66b500B80070F3280198CEBA2Af830D8e0ab",
			Value:                "0.001",
			Nonce:                fmt.Sprintf("%d", time.Now().UnixNano()),
			GasLimit:             "63262",
			GasPrice:             "0",
			MaxFeePerGas:         "2.475",
			MaxPriorityFeePerGas: "2.475",
			Chain:                "HECO",
			TransactionType:      "transfer",
		},
	})
	fmt.Println(resp, err)
}
