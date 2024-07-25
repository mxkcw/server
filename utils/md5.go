package utils

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

var mutex sync.Mutex

// MD5EncryptionGo MD5加密
func MD5EncryptionGo(str string) string {
	mutex.Lock()         // 在函数开始处加锁
	defer mutex.Unlock() // 确保函数退出时解锁
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
