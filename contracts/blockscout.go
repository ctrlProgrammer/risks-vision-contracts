package contracts

type BlockscoutAddress struct {
	IsContract bool   `json:"is_contract"`
	IsScam     bool   `json:"is_scam"`
	IsVerified bool   `json:"is_verified"`
	Name       string `json:"name"`
	Reputation string `json:"reputation"`
}

type BlockscoutToken struct {
	AddressHash          string      `json:"address_hash"`
	CirculatingMarketCap interface{} `json:"circulating_market_cap"`
	Decimals             string      `json:"decimals"`
	ExchangeRate         interface{} `json:"exchange_rate"`
	HoldersCount         string      `json:"holders_count"`
	IconURL              string      `json:"icon_url"`
	Name                 string      `json:"name"`
	Reputation           string      `json:"reputation"`
	Symbol               string      `json:"symbol"`
	TotalSupply          string      `json:"total_supply"`
	Type                 string      `json:"type"`
	Volume24h            interface{} `json:"volume_24h"`
}

type BlockscoutAddressInfo struct {
	BlockNumberBalanceUpdatedAt int64            `json:"block_number_balance_updated_at"`
	CoinBalance                 string           `json:"coin_balance"`
	CreationStatus              string           `json:"creation_status"`
	CreationTransactionHash     string           `json:"creation_transaction_hash"`
	CreatorAddressHash          string           `json:"creator_address_hash"`
	ENSDomainName               *string          `json:"ens_domain_name"`
	ExchangeRate                string           `json:"exchange_rate"`
	HasBeaconChainWithdrawals   bool             `json:"has_beacon_chain_withdrawals"`
	HasLogs                     bool             `json:"has_logs"`
	HasTokenTransfers           bool             `json:"has_token_transfers"`
	HasTokens                   bool             `json:"has_tokens"`
	HasValidatedBlocks          bool             `json:"has_validated_blocks"`
	Hash                        string           `json:"hash"`
	IsContract                  bool             `json:"is_contract"`
	IsScam                      bool             `json:"is_scam"`
	IsVerified                  bool             `json:"is_verified"`
	Metadata                    interface{}      `json:"metadata"`
	Name                        string           `json:"name"`
	ProxyType                   *string          `json:"proxy_type"`
	PublicTags                  []string         `json:"public_tags"`
	Reputation                  string           `json:"reputation"`
	Token                       *BlockscoutToken `json:"token"`
}

type BlockscoutTokenCounters struct {
	HoldersCount   string `json:"token_holders_count"`
	TransfersCount string `json:"transfers_count"`
}

type BlockscoutHoldersItem struct {
	Address BlockscoutAddress `json:"address"`
	Value   string            `json:"value"`
	TokenId string            `json:"token_id"`
}

type BlockscoutHolders struct {
	Items []BlockscoutHoldersItem `json:"items"`
}

type BlockscoutHoldersAnalysis struct {
	TotalSupply          string `json:"total_supply"`
	HoldersCount         string `json:"holders_count"`
	TopHoldersSupply     string `json:"top_holders_supply"`
	TopHoldersPercentage string `json:"top_holders_percentage"`
}
