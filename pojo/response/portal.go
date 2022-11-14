package response

type BasePagination struct {
	Total        int64 `json:"total"`
	CurrentPage  int   `json:"currentPage"`
	PrePageCount int   `json:"prePageCount"`
}

type UserLogin struct {
	JwtToken      string `json:"jwt_token"`
	UserName      string `json:"user_name"`
	AvatarAddress string `json:"avatar_address"`
}

type RegisterUser struct {
}

type PasswordUpdate struct {
}

type GetNonce struct {
	Nonce string `json:"nonce"`
}

type Test struct {
}
