package apisdk

type RespError struct {
	Ok      bool   `json:"ok"`
	ErrMsg  string `json:"err_msg"`
	ErrCode any    `json:"err_code"`
}

type RespApprovals struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Data  []struct {
		Cursor            int    `json:"cursor"`
		RecordId          string `json:"record_id"`
		OriginRecordId    string `json:"origin_record_id"`
		MyStatus          string `json:"my_status"`
		Status            string `json:"status"`
		CompanyWalletId   string `json:"company_wallet_id"`
		HDWalletID        string `json:"hd_wallet_id"`
		WalletName        string `json:"wallet_name"`
		ApproveType       string `json:"approve_type"`
		ActionType        string `json:"action_type"`
		OriginUserUuid    string `json:"origin_user_uuid"`
		OriginUserAccount string `json:"origin_user_account"`
		CreateTime        string `json:"create_time"`
		Note              string `json:"note"`
		ExtraData         struct {
			Txinfo struct {
				Eip1559              bool   `json:"eip1559"`
				To                   string `json:"to"`
				TokenAddress         string `json:"tokenAddress"`
				From                 string `json:"from"`
				Value                string `json:"value"`
				Nonce                string `json:"nonce"`
				GasLimit             string `json:"gasLimit"`
				GasPrice             string `json:"gasPrice"`
				MaxFeePerGas         string `json:"maxFeePerGas"`
				MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
				Unit                 string `json:"unit"`
				ShortHost            string `json:"shortHost"`
				Chain                string `json:"chain"`
				IsNative             bool   `json:"isNative"`
				Token                any
				Data                 string `json:"data"`
				GoogleRecordID       string `json:"google_record_id"`
				TransactionType      string `json:"transaction_type"`
				IsEnterpriseWallet   bool   `json:"isEnterpriseWallet"`
				Amount               string `json:"amount"`
				Symbol               string `json:"symbol"`
				TotalGas             string `json:"totalGas"`
				NativeCurrency       string `json:"nativeCurrency"`
				TXInput              struct {
					ChainID               string `json:"chainId"`
					TXMode                int    `json:"txMode"`
					MaxInclusionFeePerGas string `json:"maxInclusionFeePerGas"`
					MaxFeePerGas          string `json:"maxFeePerGas"`
					GasLimit              string `json:"gasLimit"`
					Nonce                 string `json:"nonce"`
					Transaction           struct {
						Transfer struct {
							Amount string `json:"amount"`
						} `json:"transfer"`
					} `json:"transaction"`
					ToAddress string `json:"toAddress"`
				} `json:"txinput"`
				GetPriceKey string `json:"getPriceKey"`
			} `json:"txinfo"`
		} `json:"extra_data"`
	} `json:"data"`
}

type RespBalance struct {
	Total  int `json:"total"`
	Assets []struct {
		TokenAddress   string `json:"token_address"`
		Balance        string `json:"balance"`
		Decimal        int    `json:"decimal"`
		Symbol         string `json:"symbol"`
		CurrencyAmount string `json:"currency_amount"`
		Price          string `json:"price"`
	} `json:"assets"`
}

type RespTxHistory struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Data  []struct {
		RecordId                      string `json:"record_id"`
		BlockHash                     string `json:"block_hash"`
		BlockNumber                   string `json:"block_number"`
		ChainName                     string `json:"chain_name"`
		TransactionHash               string `json:"transaction_hash"`
		TransactionType               string `json:"transaction_type"`
		ContractAddress               string `json:"contract_address"`
		Status                        string `json:"status"`
		TxTime                        int    `json:"tx_time"`
		From                          string `json:"from"`
		To                            string `json:"to"`
		Amount                        string `json:"amount"`
		FeeAmount                     string `json:"fee_amount"`
		FeeData                       string `json:"fee_data"`
		DappData                      string `json:"dapp_data"`
		EventLog                      string `json:"event_log"`
		ParseData                     string `json:"parse_data"`
		ClientData                    string `json:"client_data"`
		OperateType                   string `json:"operate_type"`
		TransferId                    string `json:"transfer_id"`
		ApproveLogicData              string `json:"approve_logic_data"`
		CompanyWalletApproveLogicData struct {
			RecordId          string `json:"record_id"`
			ActionType        string `json:"action_type"`
			Status            string `json:"status"`
			AgreeCount        int    `json:"agree_count"`
			TotalCount        int    `json:"total_count"`
			ApproveTransferID string `json:"approve_transfer_id"`
		} `json:"company_wallet_approve_logic_data"`
		Memo string `json:"memo"`
	} `json:"data"`
}

type RespAddCustomToken struct {
	Message string `json:"message"`
}

type RespTransactionHistory struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Data  []struct {
		RecordId                      string `json:"record_id"`
		BlockHash                     string `json:"block_hash"`
		BlockNumber                   string `json:"block_number"`
		ChainName                     string `json:"chain_name"`
		TransactionHash               string `json:"transaction_hash"`
		TransactionType               string `json:"transaction_type"`
		ContractAddress               string `json:"contract_address"`
		Status                        string `json:"status"`
		TxTime                        int    `json:"tx_time"`
		From                          string `json:"from"`
		To                            string `json:"to"`
		Amount                        string `json:"amount"`
		FeeAmount                     string `json:"fee_amount"`
		FeeData                       string `json:"fee_data"`
		DappData                      string `json:"dapp_data"`
		EventLog                      string `json:"event_log"`
		ParseData                     string `json:"parse_data"`
		ClientData                    string `json:"client_data"`
		OperateType                   string `json:"operate_type"`
		TransferId                    string `json:"transfer_id"`
		ApproveLogicData              string `json:"approve_logic_data"`
		CompanyWalletApproveLogicData string `json:"company_wallet_approve_logic_data"`
	} `json:"data"`
}

type RespNewApproval struct {
	Ok   bool `json:"ok"`
	Data struct {
		OriginRecordId string `json:"origin_record_id"`
		RecordId       string `json:"record_id"`
		IsAutoPass     bool   `json:"is_auto_pass"`
	} `json:"data"`
}

type RespAgreeApproval struct {
	Ok   bool `json:"ok"`
	Data struct {
		RecordId       string `json:"record_id"`
		OriginRecordId string `json:"origin_record_id"`
		Status         string `json:"status"`
	} `json:"data"`
}

type RespRejectApproval struct {
	RespAgreeApproval
}

type RespGetCompanyWalletInfo struct {
	Ok   bool `json:"ok"`
	Data struct {
		CompanyWalletInfo struct {
			CompanyWalletID  string `json:"company_wallet_id"`
			WalletName       string `json:"wallet_name"`
			TeamAdminMembers int    `json:"team_admin_members"`
			PolicyMembers    int    `json:"policy_members"`
		} `json:"company_wallet_info"`
		Role        string `json:"role"`
		AddressList []struct {
			Chain   string `json:"chain"`
			Address string `json:"address"`
		} `json:"address_list"`
		HDWalletList []struct {
			HDWalletID string `json:"hd_wallet_id"`
			WalletName string `json:"wallet_name"`
		} `json:"hd_wallet_list"`
	} `json:"data"`
}

type RespGetCompanyWalletHDWalletAddress struct {
	Ok   bool `json:"ok"`
	Data struct {
		AddressList []struct {
			Chain   string `json:"chain"`
			Address string `json:"address"`
		} `json:"address_list"`
	}
}
