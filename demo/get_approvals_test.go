package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetApprovals(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.GetApprovals(&apisdk.ParamGetApprovals{
		Page:   1,
		Limit:  20,
		Status: "ING",
	})
	fmt.Println(resp, err)
}

func TestGetApprovalsV2(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.GetApprovalsV2(&apisdk.ParamGetApprovalsV2{
		Page:     1,
		Limit:    20,
		ListType: "sponsor",
	})
	fmt.Println(resp, err)
}
