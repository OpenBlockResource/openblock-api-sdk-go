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

func TestNewSolTransactionSignatureApproval(t *testing.T) {
	client := apisdk.NewClient("361bdf3a1e0640979a3e2240c3361609", "U3IXmFgR848q5XA1vQVgdvW1Z69UvQXD", 10*time.Second)
	// client.CompanyWallet.Host = "hkdev8back.openblock.vip"
	resp, err := client.CompanyWallet.NewApproval(&apisdk.ParamNewApproval{
		Action: "TRANSACTION_SIGNATURE",
		TXInfo: apisdk.TXInfo{
			Chain:  "Solana",
			Type:   "dapp",
			Method: "solana_signMessage",
			Msg: &apisdk.Msg{
				SignMsg:     "7b226e6f6e6365223a223631313061393939636634323962303632343939383632633538363934613064222c2274696d657374616d70223a313736343735303235373335337d",
				Message:     "{\"nonce\":\"6110a999cf429b062499862c58694a0d\",\"timestamp\":1764750257353}",
				OriginalMsg: "7b226e6f6e6365223a223631313061393939636634323962303632343939383632633538363934613064222c2274696d657374616d70223a313736343735303235373335337d",
			},
			DappInfo: &apisdk.DappInfo{
				Origin:   "https://app.unitas.so",
				Href:     "https://app.unitas.so/earn",
				PortName: "https://app.unitas.so1764750194493",
				Icon:     "",
				Provider: "solana",
				DappName: "",
			},
			OriginAddress: "HM7LoUtJgRfChJv1m9LJH2b4vJ6W6Qn7on3sUptCQYNT",
		},
	})
	fmt.Println(resp, err)
}
