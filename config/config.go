package config

import (
	"github.com/mxkcw/windIneLog"
	"github.com/mxkcw/windIneLog/windIne_log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	ProjectName        = "server"
	ProjectVersion     = "v0.0.1"
	ProjectDescription = "social"
	HTTPRequestTimeOut = 30
	RunAsAdress        = "0.0.0.0"
	ConnectTimeout     = "120s"
)

var (
	Config         *Conf
	CurrentRunMode = windIne.RunModeDebug
	CurrentPKGMode PKGMode
)

type PKGMode int

const (
	PKGModeWithUnknown PKGMode = iota
	PKGModeWithManage
	PKGModeWithMobile
)

var stringToPKGMode = map[string]PKGMode{
	"manage":  PKGModeWithManage,
	"webSite": PKGModeWithMobile,
	"Unknown": PKGModeWithUnknown,
}

func (art PKGMode) String() string {
	switch art {
	case PKGModeWithManage:
		return "manage"
	case PKGModeWithMobile:
		return "webSite"
	case PKGModeWithUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

type Conf struct {
	System *System `yaml:"system"`
	MySql  *MySql  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
}

type System struct {
	AppEnv   string `yaml:"appEnv"`
	Domain   string `yaml:"domain"`
	Version  string `yaml:"version"`
	HttpPort string `yaml:"httpPort"`
	Host     string `yaml:"host"`
	Jwt      string `yaml:"jwt"`
}

type MySql struct {
	DbHost   string `yaml:"dbHost"`
	DbPort   int    `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Redis struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPwd"`
	RedisDbName   int    `yaml:"redisDbName"`
	RedisNetwork  string `yaml:"redisNetwork"`
}

type FileUpload struct {
	PathFormal            string `yaml:"pathFormal"`
	PathAvatar            string `yaml:"pathAvatar"`
	PathAvatarGroup       string `yaml:"pathAvatarGroup"`
	PathImageUnread       string `yaml:"pathImageUnread"`
	PathImageRead         string `yaml:"pathImageRead"`
	PathImageTimeout      string `yaml:"pathImageTimeout"`
	PathVoiceUnread       string `yaml:"pathVoiceUnread"`
	PathVoiceRead         string `yaml:"pathVoiceRead"`
	PathVoiceTimeout      string `yaml:"pathVoiceTimeout"`
	PathVoiceMp3          string `yaml:"pathVoiceMp3"`
	PathProfilePhoto      string `yaml:"pathProfilePhoto"`
	PathProfilePvoice     string `yaml:"pathProfilePvoice"`
	PathShortvideoNormal  string `yaml:"pathShortvideoNormal"`
	PathShortvideoTimeout string `yaml:"pathShortvideoTimeout"`
}

func InitConfig(configName, configPath string) {
	windIne_log.LogInfof("------------------%s", configPath)
	workDir, _ := os.Getwd()
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(workDir, configPath))
	windIne_log.LogInfof("====================%s", filepath.Join(workDir, configPath))
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
