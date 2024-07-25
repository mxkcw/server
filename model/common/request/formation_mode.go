package request

type FormationMode struct {
	UtmSource    string `form:"utmSource"`
	UtmMedium    string `form:"utmMedium"`
	UtmCampaign  string `form:"utmCampaign"`
	CurrencyCode string `form:"currencyCode"`
	PageType     string `form:"pageType"`
}
