package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetBalance(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.GetBalance(&apisdk.ParamGetBalance{
		ChainName: "Polygon",
		Currency:  "USD",
		Page:      1,
		Limit:     20,
	})
	fmt.Println("resp: ", resp)
	fmt.Println("err: ", err)
}
