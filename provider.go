package gopaccountsdk

import (
	"fmt"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// provider name
const (
	ProviderFacebook = "Facebook"
	ProviderTwitter  = "Twitter"
	ProviderGitHub   = "GitHub"
	ProviderWeChat   = "WeChat"
)

// Provider binding status
type ProviderBindsStatus struct {
	Twitter  bool `json:"Twitter"`
	Github   bool `json:"Github"`
	Facebook bool `json:"Facebook"`
	Wechat   bool `json:"WeChat"`
}

func (c *GopAccountClient) GetProviderWeChatQRCode(providerId string) (qrData, tickt string, err error) {

	urlf := casdoorsdk.GetUrl("get-qrcode", map[string]string{
		"id":           fmt.Sprintf("%s/%s", c.Claim.Owner, providerId),
		"access_token": c.AccessToken,
	})
	res, err := casdoorsdk.DoGetResponse(urlf)
	if err != nil || res.Status != "ok" {
		return "", "", fmt.Errorf("get wechat login qrcode failed")
	}
	return res.Data.(string), res.Data2.(string), nil
}

// get current user application providers
func (c *GopAccountClient) GetProviders() []*casdoorsdk.ProviderItem {
	user, err := c.GetUser()
	if err != nil {
		return make([]*casdoorsdk.ProviderItem, 0)
	}
	ps, err := casdoorsdk.GetApplication(user.SignupApplication)
	if err != nil {
		return make([]*casdoorsdk.ProviderItem, 0)
	}
	return ps.Providers
}
