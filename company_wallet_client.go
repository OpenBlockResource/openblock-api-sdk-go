package apisdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CompanyWalletClient struct {
	Host       string
	APIKey     string
	APISecret  string
	httpClient *http.Client
}

func NewCompanyWalletClient(APIKey, APISecret string, timeout time.Duration) *CompanyWalletClient {
	return &CompanyWalletClient{
		Host:      "hkdev4back.openblock.vip",
		APIKey:    APIKey,
		APISecret: APISecret,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// GetApprovals gets the lists of approvals
func (c *CompanyWalletClient) GetApprovals(params *ParamGetApprovals) (*RespApprovals, error) {
	var inparams = map[string]any{}
	if params.Page != 0 {
		inparams["page"] = params.Page
	}
	if params.Limit != 0 {
		inparams["limit"] = params.Limit
	}
	if params.Status != "" {
		inparams["status"] = params.Status
	}

	var ret = &RespApprovals{}
	err := c.invokeAPI(http.MethodGet, "/openapi/company_wallet/approvals/", inparams, ret)
	return ret, err
}

func (c *CompanyWalletClient) GetBalance(params *ParamGetBalance) (*RespBalance, error) {
	var inparams = map[string]any{}
	if params.ChainName == "" {
		return nil, fmt.Errorf("ChainName is required")
	} else {
		inparams["chain_name"] = params.ChainName
	}
	if params.Currency != "" {
		inparams["currency"] = params.Currency
	}
	if params.HDWalletID != "" {
		inparams["hd_wallet_id"] = params.HDWalletID
	}
	if params.Page != 0 {
		inparams["page"] = params.Page
	}
	if params.Limit != 0 {
		inparams["limit"] = params.Limit
	}

	var ret = &RespBalance{}
	err := c.invokeAPI(http.MethodGet, "/openapi/company_wallet/balance/", inparams, ret)
	return ret, err
}

func (c *CompanyWalletClient) GetTxHistory(params *ParamGetTxHistory) (*RespTxHistory, error) {
	var inparams = map[string]any{}
	if params.ChainName == "" {
		return nil, fmt.Errorf("ChainName is required")
	} else {
		inparams["chain_name"] = params.ChainName
	}
	if params.Page != 0 {
		inparams["page"] = params.Page
	}
	if params.Limit != 0 {
		inparams["limit"] = params.Limit
	}
	if params.OrderBy != "" {
		inparams["order_by"] = params.OrderBy
	}
	if params.Asc != 0 {
		inparams["asc"] = params.Asc
	}
	if params.HDWalletID != "" {
		inparams["hd_wallet_id"] = params.HDWalletID
	}

	var ret = &RespTxHistory{}
	err := c.invokeAPI(http.MethodGet, "/openapi/company_wallet/tx_history/", inparams, ret)
	return ret, err
}

func (c *CompanyWalletClient) AddCustomToken(params *ParamAddCustomToken) (*RespAddCustomToken, error) {
	var inparams = map[string]any{}
	if params.ChainName == "" {
		return nil, fmt.Errorf("ChainName is required")
	} else {
		inparams["chain_name"] = params.ChainName
	}
	if params.HDWalletID != "" {
		inparams["hd_wallet_id"] = params.HDWalletID
	}
	if params.SyncWallets != 0 {
		inparams["sync_wallets"] = params.SyncWallets
	}
	if params.TokenData.Address == "" {
		return nil, fmt.Errorf("TokenData.Address is required")
	}
	if params.TokenData.Symbol == "" {
		return nil, fmt.Errorf("TokenData.Symbol is required")
	}
	if params.TokenData.Decimals == 0 {
		return nil, fmt.Errorf("TokenData.Decimals is required")
	}

	tokenDataString, err := json.Marshal(map[string]any{
		"address":  params.TokenData.Address,
		"symbol":   params.TokenData.Symbol,
		"decimals": params.TokenData.Decimals,
	})
	if err != nil {
		return nil, err
	}
	inparams["token_data"] = string(tokenDataString)

	var ret = &RespAddCustomToken{}
	err = c.invokeAPI(http.MethodPost, "/openapi/company_wallet/custom_token/", inparams, ret)
	return ret, err
}

func (c *CompanyWalletClient) NewApproval(params *ParamNewApproval) (*RespNewApproval, error) {
	var inparams = map[string]any{}
	if params.Action == "" {
		return nil, fmt.Errorf("action is required")
	} else {
		inparams["action"] = params.Action
	}
	if params.HDWalletID != "" {
		inparams["hd_wallet_id"] = params.HDWalletID
	}
	if params.TXInfo.TransactionType == "" {
		return nil, fmt.Errorf("TXInfo.TransactionType is required")
	}
	if params.TXInfo.Chain == "" {
		return nil, fmt.Errorf("TXInfo.Chain is required")
	}
	if params.TXInfo.From == "" {
		return nil, fmt.Errorf("TXInfo.From is required")
	}
	txinfoString, err := json.Marshal(params.TXInfo)
	if err != nil {
		return nil, err
	}
	inparams["txinfo"] = string(txinfoString)

	var ret = &RespNewApproval{}
	err = c.invokeAPI(http.MethodPost, "/openapi/company_wallet/approval/new/", inparams, ret)
	return ret, err
}

func (c *CompanyWalletClient) AgreeApproval(params *ParamAgreeApproval) (*RespAgreeApproval, error) {
	var inparams = map[string]any{}
	if params.RecordID == "" {
		return nil, fmt.Errorf("RecordID is required")
	} else {
		inparams["record_id"] = params.RecordID
	}
	if params.Agree == "" {
		return nil, fmt.Errorf("Agree is required")
	} else {
		inparams["agree"] = params.Agree
	}

	var ret = &RespAgreeApproval{}
	err := c.invokeAPI(http.MethodPost, "/openapi/company_wallet/approval/agree/", inparams, ret)
	return ret, err
}

func (c *CompanyWalletClient) GetCompanyWalletInfo() (*RespGetCompanyWalletInfo, error) {
	var inparams = map[string]any{}
	var ret = &RespGetCompanyWalletInfo{}
	err := c.invokeAPI(http.MethodGet, "/openapi/company_wallet/info/", inparams, ret)
	return ret, err
}

func (c *CompanyWalletClient) GetCompanyWalletHDWalletAddress(params *ParamGetCompanyWalletHDWalletAddress) (*RespGetCompanyWalletHDWalletAddress, error) {
	var inparams = map[string]any{}
	if params.HDWalletID == "" {
		return nil, fmt.Errorf("HDWalletID is required")
	} else {
		inparams["hd_wallet_id"] = params.HDWalletID
	}
	var ret = &RespGetCompanyWalletHDWalletAddress{}
	err := c.invokeAPI(http.MethodGet, "/openapi/company_wallet/hd_wallet_address/", inparams, ret)
	return ret, err
}

func (c *CompanyWalletClient) invokeAPI(method, path string, params map[string]any, response any) error {
	bodyBytes, err := c.doRequest(method, path, params)
	if err != nil {
		return err
	}

	// Check if error occured
	var bodyMap map[string]any
	err = json.Unmarshal(bodyBytes, &bodyMap)
	if err != nil {
		return err
	}
	if _, ok := bodyMap["err_code"]; ok {
		var retErr RespError
		err = json.Unmarshal(bodyBytes, &retErr)
		if err != nil {
			return err
		}
		return fmt.Errorf("error occured: code=%s msg=%s", retErr.ErrCode, retErr.ErrMsg)
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return err
	}
	return nil
}

// doRequest returns response body and error
func (c *CompanyWalletClient) doRequest(method string, path string, params map[string]any) ([]byte, error) {
	// Create HTTP request
	var req *http.Request
	var err error
	switch method {
	case http.MethodGet:
		params = c.genParams(params)
		u := url.URL{
			Scheme:   "https",
			Host:     c.Host,
			Path:     path,
			RawQuery: mapToValues(params).Encode(),
		}
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}
	case http.MethodPost:
		params = c.genParams(params)
		u := url.URL{
			Scheme: "https",
			Host:   c.Host,
			Path:   path,
		}
		jsonBody, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, u.String(), bytes.NewReader(jsonBody))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
	default:
		return nil, fmt.Errorf("invalid method: %s", method)
	}

	// 发送请求并获取响应
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	return bodyBytes, err
}

func (c *CompanyWalletClient) genParams(params map[string]any) map[string]any {
	ret := make(map[string]any)
	for k, v := range params {
		ret[k] = v
	}

	nonce, ok := params["nonce"]
	if !ok || nonce == "" {
		uuid4 := uuid.New()
		ret["nonce"] = uuid4.String()
	}
	// Upper
	ret["nonce"] = strings.ToUpper(ret["nonce"].(string))

	ret["sign"] = c.genSign(ret)
	ret["api_key"] = c.APIKey
	return ret
}

func (c *CompanyWalletClient) genSign(params map[string]any) string {
	// order keys in alpha ascending
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var strs string
	for _, k := range keys {
		v := params[k]
		strs += fmt.Sprintf("%s=%v&", k, v)
	}
	strs += fmt.Sprintf("api_key=%s", c.APIKey)

	return strings.ToUpper(hmacSha256(c.APISecret, strs))
}

func mapToValues(m map[string]any) url.Values {
	ret := url.Values{}
	for k, v := range m {
		ret.Add(k, fmt.Sprintf("%v", v))
	}
	return ret
}

func hmacSha256(key string, message string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
