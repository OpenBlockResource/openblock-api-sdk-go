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
import (
    "time"
    apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

# Refer to https://docs.openblock.com/OpenBlock/API/Enterprise%20Wallet/#overview to get your api key and secret
client = apisdk.NewClient(
    "YOUR API KEY",
    "YOUR SECRET",
    10 * time.Second,
)

resp := client.GetBalance(
    "Polygon",
    "USD",
    1, 20
)
fmt.Printf("%+v\n", resp)
```
