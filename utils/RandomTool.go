package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 获取 map 中的随机键
func getRandomKey(m map[string]string) string {
	// 创建一个切片来存放键
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// 生成一个随机索引
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(keys))
	// 返回随机value
	return m[fmt.Sprintf("%d", randomIndex)]
}

func RandomAvatar() string {
	avatar := map[string]string{
		"1":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/10.jpeg",
		"2":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/11.jpeg",
		"3":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/12.jpeg",
		"4":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/13.jpeg",
		"5":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/14.jpeg",
		"6":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/15.jpeg",
		"7":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/16.jpeg",
		"8":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/17.jpeg",
		"9":  "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/18.jpeg",
		"10": "https://xd-video-pc-img.oss-cn-beijing.aliyuncs.com/xdclass_pro/default/head_img/19.jpeg",
	}
	return getRandomKey(avatar)
}

func RandomName() string {
	name := map[string]string{
		"1":  "小白23423",
		"2":  "小白94352",
		"3":  "小白27871",
		"4":  "小白12312",
		"5":  "小白19292",
		"6":  "小白30239",
		"8":  "小白12344",
		"9":  "小白12345",
		"10": "小白93841",
	}
	return getRandomKey(name)
}
