package common

import "time"

type SiteFormation struct {
	ID           uint64    `json:"id" gorm:"primary_key;auto_increment;not_null;comment:'id'"`
	ApiKey       string    `json:"api_key" gorm:"type:varchar(255);charset:utf8mb4;collate:utf8mb4_0900_ai_ci;comment:'apiKey and wegitID'"`
	CurrencyCode string    `json:"currency_code" gorm:"type:varchar(255);comment:'currency code'"`
	UTMSource    string    `json:"utm_source" gorm:"type:varchar(255);comment:'source'"`
	UTMMedium    string    `json:"utm_medium" gorm:"type:varchar(255);comment:'medium- video - article - other'"`
	UTMCampaign  string    `json:"utm_campaign" gorm:"type:varchar(255);comment:'campaign'"`
	URL          string    `json:"url" gorm:"type:varchar(255);comment:'url'"`
	GMTCreate    time.Time `json:"gmt_create" gorm:"type:datetime;comment:'gmt_create'"`
	GMTModified  time.Time `json:"gmt_modified" gorm:"type:datetime;comment:'gmt_modified'"`
	PageType     string    `json:"page_type" gorm:"type:varchar(255);comment:'page_type'"`
	State        int64     `json:"state" gorm:"type:int(2);comment:'1:no send 2:send'"`
}

func (SiteFormation) TableName() string {
	return "site_formation"
}
