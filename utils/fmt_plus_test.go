package utils

import (
	"fmt"
	"github.com/mxkcw/windIneLog/windIne_log"
	"server/middleware"
	"testing"
)

func TestRandomNumber(t *testing.T) {

	number := RandomNumber(4)
	windIne_log.LogInfof("%s", number)
}

func TestRandomName(t *testing.T) {
	name := RandomName()
	avatar := RandomAvatar()
	windIne_log.LogInfof("%s", name)
	windIne_log.LogInfof("%s", avatar)
}

func TestToken(f *testing.T) {
	token, err := middleware.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjY4MTQ0MjIiLCJQaG9uZSI6IjE4MTgwNjc0MTM5IiwiQXZhdGFyIjoiaHR0cHM6Ly94ZC12aWRlby1wYy1pbWcub3NzLWNuLWJlaWppbmcuYWxpeXVuY3MuY29tL3hkY2xhc3NfcHJvL2RlZmF1bHQvaGVhZF9pbWcvMTguanBlZyIsIk5hbWUiOiLlsI_nmb0xOTI5MiIsImV4cCI6MTcyMDUzMDM0MywiaXNzIjoiYWRtaW4ifQ.pfBXcM4XiaZpwJRFCArZVm9vBfrCUUWWcrUSxdZDmzs ")

	windIne_log.LogInfof("%s", err)
	windIne_log.LogInfof("%s,%s", token.ID, token.Phone)
}

func TestMD5EncryptionGo(t *testing.T) {
	encryptionGo := MD5EncryptionGo("qwe1230.")
	fmt.Println("encryptionGo", encryptionGo)
}
