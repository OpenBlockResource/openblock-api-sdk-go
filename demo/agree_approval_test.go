package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestAgreeApproval(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	resp, err := client.CompanyWallet.AgreeApproval(&apisdk.ParamAgreeApproval{
		RecordID: "recordId",
	})
	fmt.Println(resp, err)
}
