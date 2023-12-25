package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestAddCustomToken(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.AddCustomToken(&apisdk.ParamAddCustomToken{
		ChainName: "Polygon",
		TokenData: apisdk.TokenData{
			Address:  "0xABCD",
			Symbol:   "USDT",
			Decimals: 6,
		},
	})
	fmt.Println(resp, err)
}
