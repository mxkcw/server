package request

import "time"

type UserLogin struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

type AddVistRecord struct {
	DeviceType  string `json:"device_type"`
	Region      string `json:"region"`
	Referer     string `json:"referer"`
	UtmSource   string `json:"utm_source"`
	UtmMedium   string `json:"utm_medium"`
	UtmCampaign string `json:"utm_campaign"`
	ApiKey      string `json:"api_key"`
}

type Record struct {
	AddVistRecord
	Everyday time.Time
}
