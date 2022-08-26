package edx

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type NewToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type AuthUseCase interface {
	RefreshToken() (err error)
	GetWithAuth(url string) (resBody []byte, err error)
	PostWithAuth(url string, params map[string]interface{}) (respBody []byte, err error)
}
