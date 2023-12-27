# Golang SDK for Openblock OpenAPI

# API Documentation
- [Official documentation](https://docs.openblock.com/zh-Hans/OpenBlock/API/Enterprise%20Wallet/)

# Installation

```bash
go get github.com/OpenBlockResource/openblock-api-sdk-go
```

# Usage
> Consider `/openapi/company_wallet/balance/` as a reference, the full code can be located in the demo/api_demo directory.
* Get api key and secret

https://docs.openblock.com/OpenBlock/API/Enterprise%20Wallet/#overview

* Get your balance
```golang
package main

import (
    "fmt"
    apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
    "time"
)

func main() {
    client := apisdk.NewClient(
        "API KEY",
        "API Secret",
        10*time.Second,
    )

    resp, err := client.CompanyWallet.GetBalance(&apisdk.ParamGetBalance{
        ChainName: "Polygon",
        Currency:  "USD",
        Page:      0,
        Limit:     20,
    })

    fmt.Printf("%+v\n", resp)
    fmt.Printf("err: %+v\n", err)
}
```
