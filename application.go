package gopaccountsdk

import (
	"log"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// 获取当前应用
func GetApplicationInfo(appName string) (*casdoorsdk.Application, error) {
	a2, err := casdoorsdk.GetApplication(appName)
	if err != nil {
		log.Println(err)
	}
	return a2, err
}
