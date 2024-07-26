package utils

import (
	"github.com/mssola/user_agent"
)

func DeviceType(uaString string) string {
	ua := user_agent.New(uaString)
	if ua.Mobile() {
		return "Mobile"
	}
	return "PC"
}
