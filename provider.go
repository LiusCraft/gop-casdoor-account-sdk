package gopaccountsdk

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
