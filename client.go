package gopaccountsdk

import "github.com/casdoor/casdoor-go-sdk/casdoorsdk"

// Casdoor user session
type GopAccountClient struct {
	AccessToken string
	Claim       *casdoorsdk.Claims
}

// Obtain GopAccountClient through token

func GetClient(token string) (*GopAccountClient, error) {
	claim, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		return nil, err
	}
	return &GopAccountClient{
		Claim:       claim,
		AccessToken: token,
	}, nil
}
