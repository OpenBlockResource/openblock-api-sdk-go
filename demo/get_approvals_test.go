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
	client := apisdk.NewClient("361bdf3a1e0640979a3e2240c3361609", "U3IXmFgR848q5XA1vQVgdvW1Z69UvQXD", 10*time.Second)
	client.CompanyWallet.Host = "hkdev8back.openblock.vip"
	resp, err := client.CompanyWallet.GetApprovalsV2(&apisdk.ParamGetApprovalsV2{
		Page:     1,
		Limit:    20,
		ListType: "sponsor",
		RecordID: "586d043897734256b3fa402b08d464bd",
	})
	fmt.Printf("resp: %+v,err: %v", resp, err)
}
