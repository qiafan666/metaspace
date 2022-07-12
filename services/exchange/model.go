package exchange

import (
	"time"
)

type CoinMarketRequest struct {
	Status Status `json:"status"`
	Data   []Data `json:"data"`
}
type Status struct {
	Timestamp    time.Time   `json:"timestamp"`
	ErrorCode    int         `json:"error_code"`
	ErrorMessage interface{} `json:"error_message"`
	Elapsed      int         `json:"elapsed"`
	CreditCount  int         `json:"credit_count"`
	Notice       interface{} `json:"notice"`
}

type Data struct {
	Id          int             `json:"id"`
	Symbol      string          `json:"symbol"`
	Name        string          `json:"name"`
	Amount      int             `json:"amount"`
	LastUpdated time.Time       `json:"last_updated"`
	Quote       map[string]Coin `json:"quote"`
}

type Coin struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}
