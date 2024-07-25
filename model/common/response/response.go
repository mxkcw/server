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

func GetResCodeDataError(msg string, data interface{}, isDebug bool, code int64) *ResData {
	aErr := GetDefaultRespBase()
	aErr.Code = code
	aErr.Msg = msg
	if data != nil {
		if !isDebug {
			aErr.Data = data
		}
	}
	return aErr
}

func GetResCodeDataSuccess(msg string, data interface{}, isDebug bool) *ResData {
	aErr := GetDefaultRespBase()
	aErr.Code = SUCCESS
	aErr.Msg = msg
	if data != nil {
		if !isDebug {
			aErr.Data = data
		}
	}

	return aErr
}
