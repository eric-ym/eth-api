package service

type BlocksResponse struct {
	BlockID string        `json:"blockID"`
	Height  int64         `json:"height"`
	Time    int64         `json:"time"`
	Data    []MessageInfo `json:"data"`
}

type MessageInfo struct {
	Cid             string `json:"cid"`
	StorageProvider string `json:"StorageProvider"`
	Message         int    `json:"message"`
	Reward          string `json:"reward"`
}

// BlockInfoResponse  block info
type BlockInfoResponse struct {
	Cid               string `json:"Cid"`
	Height            int64  `json:"height"`
	Time              int64  `json:"time"`
	Message           int    `json:"message"`
	StorageProvider   string `json:"StorageProvider"`
	WinCount          int64  `json:"winCount"`
	ParentCid         string `json:"parentCid"`
	ParentWeight      int64  `json:"parentWeight"`
	ParentBaseFeeRate string `json:"parentBaseFeeRate"`
	Ticket            string `json:"ticket"`
	StateRoot         string `json:"stateRoot"`
	Reward            string `json:"reward"`
}

type BlockReward struct {
	Total int64 `json:"total"`
	Block int64 `json:"block"`
	Fee   int64 `json:"fee"`
}

type MessageResponse struct {
	Id     string `json:"id"`
	Height int64  `json:"height"`
	Block  string `json:"block"`
	Time   int64  `json:"time"`
	From   string `json:"from"`
	To     string `json:"to"`
	Value  string `json:"value"`
	Status string `json:"status"`
	Method string `json:"method"`
}

type MessageData struct {
	Total int               `json:"total"`
	Page  int               `json:"page"`
	Limit int               `json:"limit"`
	Data  []MessageResponse `json:"data"`
}

// MessageDetail 消息详情
type MessageDetail struct {
	Version    string `json:"version"`
	Nonce      string `json:"nonce"`
	GasFeeCap  string `json:"gasFeeCap"`
	GasPremium string `json:"gasPremium"`
	GasLimit   string `json:"gasLimit"`
	GasUsed    string `json:"gasUsed"`
	BaseFee    string `json:"baseFee"`
	GasFee     string `json:"gasFee"`
	Params     string `json:"params"`
	Return     string `json:"return"`
}

type MessageDetailTransaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AddressData struct {
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Data  []AddressInfo `json:"data"`
}

type AddressInfo struct {
	Rank           int     `json:"rank"`
	Account        string  `json:"account"`
	Balance        float64 `json:"balance"`
	Percentage     string  `json:"percentage"`
	Type           int     `json:"type"`
	LastUpdate     int64   `json:"lastUpdate"`
	RecentTransfer int64   `json:"recentTransfer"`
}

type AddressDetailInfo struct {
	Address    string `json:"address"`
	AccountID  int64  `json:"accountID"`
	Balance    string `json:"balance"`
	Message    int    `json:"message"`
	Nounce     string `json:"nounce"`
	Cid        string `json:"cid"`
	CreateTime int64  `json:"createTime"`
	LastUpdate int64  `json:"lastUpdate"`
}

type AddressMessageData struct {
	Total int                  `json:"total"`
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
	Data  []AddressMessageInfo `json:"data"`
}

type AddressMessageInfo struct {
	Id     string `json:"id"`
	Height int64  `json:"height"`
	Time   int64  `json:"time"`
	From   string `json:"from"`
	To     string `json:"to"`
	Value  string `json:"value"`
	Status string `json:"status"`
	Method string `json:"method"`
}

type AddressTransData struct {
	Total int                `json:"total"`
	Page  int                `json:"page"`
	Limit int                `json:"limit"`
	Data  []AddressTransInfo `json:"data"`
}

type AddressTransInfo struct {
	Time   int64  `json:"time"`
	Id     string `json:"id"`
	From   string `json:"from"`
	To     string `json:"to"`
	Value  string `json:"value"`
	Method string `json:"method"`
}

type Standard struct {
	LastHeight     int64  `json:"lastHeight"`     // 最新区块高度
	LastUpdate     int64  `json:"lastUpdate"`     // 最新区块时间
	HashPower      string `json:"hashPower"`      // 全网有效算力
	Power24        string `json:"power24"`        //24小时增长算力
	Productivity24 string `json:"productivity24"` //24小时产出效率
}

type HomeReward struct {
	Bonus24 int64 `json:"24hbonus"`
}

type HomeEChatData struct {
	Hour  int    `json:"hour"`
	Power string `json:"power"`
	Incr  string `json:"incr"`
}
