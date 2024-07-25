package utils

var (
	UserLogin      = Rules{"Phone": []string{NotEmpty()}, "Password": []string{NotEmpty()}}
	Formation      = Rules{"UtmSource": []string{NotEmpty()}, "UtmMedium": []string{NotEmpty()}, "CurrencyCode": []string{NotEmpty()}, "PageType": []string{NotEmpty()}}
	PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
