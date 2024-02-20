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
	Twitter  bool
	Github   bool
	Facebook bool
	Wechat   bool
}
