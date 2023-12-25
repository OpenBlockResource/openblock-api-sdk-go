package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetCompanyWalletInfo(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.GetCompanyWalletInfo()
	fmt.Println(resp, err)
}
