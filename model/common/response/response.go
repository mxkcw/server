package response

import (
	"time"
)

const (
	offsetCode = 8
)

type ResData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time int64       `json:"time"`
}

func GetDefaultRespBase() *ResData {
	return &ResData{
		Code: SUCCESS,
		Time: time.Now().UTC().UnixMicro(),
	}
}

func GetResCodeDataError(msg string, data interface{}, code int64) *ResData {
	aErr := GetDefaultRespBase()
	aErr.Code = code
	aErr.Msg = msg
	aErr.Data = data
	return aErr
}

func GetResCodeDataSuccess(msg string, data interface{}) *ResData {
	aErr := GetDefaultRespBase()
	aErr.Code = SUCCESS
	aErr.Msg = msg
	aErr.Data = data

	return aErr
}
