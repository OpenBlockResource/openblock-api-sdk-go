package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetTxHistory(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.GetTxHistory(&apisdk.ParamGetTxHistory{
		ChainName: "Polygon",
		Page:      1,
		Limit:     20,
		OrderBy:   "tx_time",
		Asc:       0,
	})
	fmt.Println(resp, err)
}
