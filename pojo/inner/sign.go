package inner

type SignRequest struct {
	ApiKey    string `json:"api_key"`
	Timestamp string `json:"timestamp"`
	Rand      string `json:"rand"`
	Uri       string `json:"uri"`
	Parameter string `json:"parameter"`
}

type SignResponse struct {
	Sign string `json:"sign"`
}

type VerifySignRequest struct {
	Sign      string `json:"sign"`
	ApiKey    string `json:"api_key"`
	Timestamp string `json:"timestamp"`
	Rand      string `json:"rand"`
	Uri       string `json:"uri"`
	Parameter string `json:"parameter"`
}
type VerifySignResponse struct {
	Result bool `json:"result"`
}
