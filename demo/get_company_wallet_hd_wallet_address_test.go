package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetCompanyWalletHDWalletAddress(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.GetCompanyWalletHDWalletAddress(&apisdk.ParamGetCompanyWalletHDWalletAddress{
		HDWalletID: "HDWalletID",
	})
	fmt.Println(resp, err)
}
