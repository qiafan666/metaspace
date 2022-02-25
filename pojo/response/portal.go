package response

type UserLogin struct {
	JwtToken string `json:"jwt_token"`
}

type RegisterUser struct {
}

type PasswordUpdate struct {
}
