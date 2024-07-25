package http

import (
	"encoding/json"
	"github.com/mxkcw/windIneLog/windIne_http"
	"github.com/mxkcw/windIneLog/windIne_log"
)

type RequestData struct {
	Appid        string `json:"appid"`
	AppSecret    string `json:"appSecret"`
	Code         string `json:"code"`
	PhoneNum     string `json:"phoneNum"`
	TemplateCode string `json:"templateCode"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

func SendMsgCode(code string, phone string) (response Response, err error) {

	requestData := RequestData{
		Code:         code,
		PhoneNum:     phone,
		Appid:        "8UVpQ7I4PqJFV8MytS",
		AppSecret:    "VQBGfh2iwgiaVxzSXP615WDluIYyp59E",
		TemplateCode: "SMS_168781429",
	}
	dataStr, err := json.Marshal(requestData)
	if err != nil {
		return response, err
	}
	windIne_log.LogInfof("dataStr:%s", dataStr)
	httpClient := windIne_http.NewHttpClient(10)
	postResult, err := httpClient.PostResult("https://api-v2.xdclass.net/send_sms", dataStr)
	err = json.Unmarshal(postResult, &response)
	return response, err
}
