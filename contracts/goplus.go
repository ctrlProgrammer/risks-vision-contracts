package contracts

type GoplusDexInfo struct {
	LiquidityType string `json:"liquidity_type"`
	Name          string `json:"name"`
	Liquidity     string `json:"liquidity"`
	Pair          string `json:"pair"`
}

type GoplusHolderInfo struct {
	Address    string `json:"address"`
	Tag        string `json:"tag"`
	IsContract int    `json:"is_contract"`
	Balance    string `json:"balance"`
	Percent    string `json:"percent"`
	IsLocked   int    `json:"is_locked"`
}

type GoplusLpHolderInfo struct {
	Address    string      `json:"address"`
	Tag        string      `json:"tag"`
	Value      interface{} `json:"value"`
	IsContract int         `json:"is_contract"`
	Balance    string      `json:"balance"`
	Percent    string      `json:"percent"`
	NFTList    interface{} `json:"NFT_list"`
	IsLocked   int         `json:"is_locked"`
}

type GoplusInCexInfo struct {
	Listed  string   `json:"listed"`
	CexList []string `json:"cex_list"`
}

type GoplusTokenData struct {
	AntiWhaleModifiable        string               `json:"anti_whale_modifiable"`
	BuyTax                     string               `json:"buy_tax"`
	CanTakeBackOwnership       string               `json:"can_take_back_ownership"`
	CannotBuy                  string               `json:"cannot_buy"`
	CannotSellAll              string               `json:"cannot_sell_all"`
	CreatorAddress             string               `json:"creator_address"`
	CreatorBalance             string               `json:"creator_balance"`
	CreatorPercent             string               `json:"creator_percent"`
	Dex                        []GoplusDexInfo      `json:"dex"`
	ExternalCall               string               `json:"external_call"`
	HiddenOwner                string               `json:"hidden_owner"`
	HolderCount                string               `json:"holder_count"`
	Holders                    []GoplusHolderInfo   `json:"holders"`
	HoneypotWithSameCreator    string               `json:"honeypot_with_same_creator"`
	IsAntiWhale                string               `json:"is_anti_whale"`
	IsBlacklisted              string               `json:"is_blacklisted"`
	IsHoneypot                 string               `json:"is_honeypot"`
	IsInCex                    GoplusInCexInfo      `json:"is_in_cex"`
	IsInDex                    string               `json:"is_in_dex"`
	IsMintable                 string               `json:"is_mintable"`
	IsOpenSource               string               `json:"is_open_source"`
	IsProxy                    string               `json:"is_proxy"`
	IsWhitelisted              string               `json:"is_whitelisted"`
	LpHolderCount              string               `json:"lp_holder_count"`
	LpHolders                  []GoplusLpHolderInfo `json:"lp_holders"`
	LpTotalSupply              string               `json:"lp_total_supply"`
	OwnerAddress               string               `json:"owner_address"`
	OwnerBalance               string               `json:"owner_balance"`
	OwnerChangeBalance         string               `json:"owner_change_balance"`
	OwnerPercent               string               `json:"owner_percent"`
	PersonalSlippageModifiable string               `json:"personal_slippage_modifiable"`
	SelfDestruct               string               `json:"selfdestruct"`
	SellTax                    string               `json:"sell_tax"`
	SlippageModifiable         string               `json:"slippage_modifiable"`
	TokenName                  string               `json:"token_name"`
	TokenSymbol                string               `json:"token_symbol"`
	TotalSupply                string               `json:"total_supply"`
	TradingCooldown            string               `json:"trading_cooldown"`
	TransferPausable           string               `json:"transfer_pausable"`
	TransferTax                string               `json:"transfer_tax"`
	TrustList                  string               `json:"trust_list"`
}

type GoplusTokenRiskResponse struct {
	Code    int                        `json:"code"`
	Message string                     `json:"message"`
	Result  map[string]GoplusTokenData `json:"result"`
}
