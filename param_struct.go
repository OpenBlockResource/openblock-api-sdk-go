package apisdk

type ParamGetApprovals struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Status string `json:"status"`
}

type ParamGetBalance struct {
	ChainName  string `json:"chain_name"`
	Currency   string `json:"currency"`
	HDWalletID string `json:"hd_wallet_id"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}

type ParamGetTxHistory struct {
	ChainName  string `json:"chain_name"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	OrderBy    string `json:"order_by"`
	Asc        int    `json:"asc"`
	HDWalletID string `json:"hd_wallet_id"`
}

type TokenData struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

type ParamAddCustomToken struct {
	ChainName   string    `json:"chain_name"`
	HDWalletID  string    `json:"hd_wallet_id"`
	SyncWallets int       `json:"sync_wallets"`
	TokenData   TokenData `json:"token_data"`
}

type DappInfo struct {
	Origin   string `json:"origin"`
	Href     string `json:"href"`
	PortName string `json:"portName"`
	Icon     string `json:"icon"`
	DappName string `json:"dappName"`
	Provider string `json:"provider,omitempty"`
}

type Utxo struct {
	Script string `json:"script"`
	Amount string `json:"amount"`
	Hash   string `json:"hash"`
	Index  string `json:"index"`
}

type TXInfo struct {
	To                   string            `json:"to,omitempty"`
	From                 string            `json:"from,omitempty"`
	TokenAddress         string            `json:"tokenAddress,omitempty"`
	Value                string            `json:"value,omitempty"`
	Nonce                string            `json:"nonce,omitempty"`
	GasLimit             string            `json:"gasLimit,omitempty"`
	GasPrice             string            `json:"gasPrice,omitempty"`
	MaxFeePerGas         string            `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string            `json:"maxPriorityFeePerGas,omitempty"`
	Chain                string            `json:"chain,omitempty"`
	TransactionType      string            `json:"transaction_type,omitempty"`
	OriginalRecordId     string            `json:"original_record_id,omitempty"`
	OperateType          string            `json:"operate_type,omitempty"`
	Data                 string            `json:"data,omitempty"`
	PretreatmentValue    string            `json:"pretreatment_value,omitempty"`
	TokenID              string            `json:"token_id,omitempty"`
	DappInfo             *DappInfo         `json:"dappInfo,omitempty"`
	Utxo                 []*Utxo           `json:"utxo,omitempty"`
	UseMaxAmount         bool              `json:"useMaxAmount,omitempty"`
	TotalGas             string            `json:"totalGas,omitempty"`
	RecentBlockHash      string            `json:"recent_blockhash"`
	ActiveTokenEnum      int               `json:"activeTokenEnum"`
	TvmParams            *TVMParams        `json:"tvmParams,omitempty"`
	Token                *TokenData        `json:"token,omitempty"`
	FeeLimit             string            `json:"feeLimit"`
	GasInfo              map[string]string `json:"gasInfo"`
	Type                 string            `json:"type,omitempty"`
	Method               string            `json:"method,omitempty"`
	Msg                  *Msg              `json:"msg,omitempty"`
	OriginAddress        string            `json:"origin_address,omitempty"`
	TxPayload            interface{}       `json:"txPayload,omitempty"`
	IsNative             bool              `json:"isNative,omitempty"`
	Eip1559              bool              `json:"eip1559,omitempty"`
	Payload              any               `json:"payload,omitempty"`
	BridgeMethod         string            `json:"bridgeMethod,omitempty"`
}

type Msg struct {
	SignMsg     string `json:"signMsg"`
	Message     string `json:"message"`
	OriginalMsg string `json:"originalMsg"`
}

type TVMParams struct {
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witnessAddress"`
	ParentHash     string `json:"parentHash"`
	Version        string `json:"version"`
	Timestamp      string `json:"timestamp"`
	Number         string `json:"number"`
}

type ParamNewApproval struct {
	Action         string `json:"action"`
	HDWalletID     string `json:"hd_wallet_id"`
	TXInfo         TXInfo `json:"txinfo"`
	Note           string `json:"note"`
	ExpiredTimeout int32  `json:"expired_timeout,omitempty"`
}

type ParamAgreeApproval struct {
	RecordID string `json:"record_id"`
	Agree    string `json:"agree"`
}

type ParamGetCompanyWalletHDWalletAddress struct {
	HDWalletID string `json:"hd_wallet_id"`
}

type ParamGetApprovalsV2 struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	ListType string `json:"list_type"`
	RecordID string `json:"record_id"`
}
