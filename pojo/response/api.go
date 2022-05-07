package response

type CreateAuthCode struct {
	AuthCode     string `json:"auth_code"`
	LoginAddress string `json:"login_address"`
	LoginUrl     string `json:"login_url"`
}
