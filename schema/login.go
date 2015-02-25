package schema

type LoginOpts struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	OauthScope string `json:"oauthScope"`
}

type LoginResult struct {
	Token string `json:"token"`
}

type LoginData struct {
	AccessToken string `json:"accessToken"`
}

type Login struct {
	Result  LoginResult `json:"result"`
	Data    LoginData   `json:"data"`
	Success bool        `json:"success"`
}
