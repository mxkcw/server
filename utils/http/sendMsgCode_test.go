package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestSendMsgCode(t *testing.T) {

	//result, err := SendMsgCode("1234", "18180674136")
	//var resp Response
	//err = json.Unmarshal(result, &resp)
	//fmt.Println(err)
	//fmt.Println(resp.Code)
	//if resp.Code == 0 {
	//	fmt.Println("code == 0", resp.Code)
	//} else {
	//	fmt.Println("code != 0", resp.Code)
	//}

	// 创建请求数据
	requestData := RequestData{
		Code:         "123411111",
		PhoneNum:     "18180674136",
		Appid:        "8UVpQ7I4PqJFV8MytS",
		AppSecret:    "n8Syf7tFYbbJuN5rC1IR7wNwFL74BFCg",
		TemplateCode: "SMS_168781429",
	}

	// 将请求数据序列化为 JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshalling request data:", err)
		return
	}
	fmt.Println("jsonData:", string(jsonData))

	// 创建一个新的请求
	url := "https://api-v2.xdclass.net/send_sms" // 替换为实际的 URL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 打印响应
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
