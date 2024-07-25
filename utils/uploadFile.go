package utils

import (
	"github.com/google/uuid"
	"github.com/mxkcw/windIneLog/windIne_log"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func Upload(file *multipart.FileHeader, subFolder string) (map[string]interface{}, error) {
	now := time.Now()
	var path string
	if subFolder == "" {
		path = now.Format("20060102")
	} else {
		path = filepath.Join(subFolder, now.Format("20060102"))
	}

	newPath := filepath.Join("filepath", path)
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		windIne_log.LogErrorf("%s：%s", "err", err.Error())
		return nil, err
	}

	// 读取文件内容
	fileData, err := file.Open()
	if err != nil {
		windIne_log.LogErrorf("%s：%s", "err", err.Error())
		return nil, err
	}
	defer fileData.Close()

	bytes, err := ioutil.ReadAll(fileData)
	if err != nil {
		windIne_log.LogErrorf("%s：%s", "err", err.Error())
		return nil, err
	}

	// 获取文件后缀
	extension := filepath.Ext(file.Filename)
	// 生成UUID文件名
	fileUUIDname := uuid.New().String()
	newFileName := fileUUIDname + extension
	newPath = filepath.Join(newPath, newFileName)

	// 保存文件
	err = ioutil.WriteFile(newPath, bytes, os.ModePerm)
	if err != nil {
		windIne_log.LogErrorf("%s：%s", "mobileConsts.ERROR", err.Error())
		return nil, err
	}

	filepath := path + "/" + newFileName

	var result = make(map[string]interface{})
	result["path"] = filepath
	result["url"] = filepath
	return result, nil
}
