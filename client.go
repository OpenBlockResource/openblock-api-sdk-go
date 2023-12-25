package apisdk

import "time"

type Client struct {
	CompanyWallet *CompanyWalletClient
}

func NewClient(APIKey, APISecret string, timeout time.Duration) *Client {
	return &Client{
		CompanyWallet: NewCompanyWalletClient(APIKey, APISecret, timeout),
	}
}
