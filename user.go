package gopaccountsdk

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// simple user
type UserSimple struct {
	Id       string
	Name     string
	Password string
	Avatar   string
}

func (c GopAccountClient) GetUser() (*casdoorsdk.User, error) {
	return casdoorsdk.GetUser(c.Claim.Name)
}

// Get Current User Provider binding status
func (c GopAccountClient) GetProviderBindStatus() ProviderBindsStatus {
	u, err := c.GetUser()
	if err != nil {
		return ProviderBindsStatus{}
	}

	return ProviderBindsStatus{
		len(u.Twitter) > 0,
		len(u.GitHub) > 0,
		len(u.Facebook) > 0,
		len(u.WeChat) > 0,
	}
}

// Unbind provider
//
// providerName: lookup provider.go
func (p GopAccountClient) UnLink(provider string) bool {
	u, err := p.GetUser()
	if err != nil {
		log.Println(err)
		return false
	}
	postData, _ := json.Marshal(map[string]interface{}{
		"providerType": provider,
		"user":         u,
	})

	var response casdoorsdk.Response
	targetUrl := casdoorsdk.GetUrl("unlink", map[string]string{
		"access_token": p.AccessToken,
	})
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(postData))

	req.Header.Add("Content-Type", "application/json")
	resultBody, _ := http.DefaultClient.Do(req)
	body, err := io.ReadAll(resultBody.Body)
	if err != nil {
		return false
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false
	}

	if response.Status != "ok" {
		log.Println(response.Msg)
		return false
	}
	return true
}

// Get UserInfo from access token
func (c GopAccountClient) GetUserSimple() *UserSimple {
	return &UserSimple{
		Name:   c.Claim.Name,
		Avatar: c.Claim.Avatar,
		Id:     c.Claim.Id,
	}
}
