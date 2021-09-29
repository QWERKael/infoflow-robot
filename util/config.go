package util

import (
	"encoding/base64"
	"github.com/QWERKael/utility-go/config"
)

var Config = &Conf{}

type Conf struct {
	Token          string `yaml:"Token"`
	EncodingAESKey string `yaml:"EncodingAESKey"`
	RobotUrl       string `yaml:"RobotUrl"`
}

func InitConfigWithYaml(configPath string) {
	err := config.ParserFromPath(configPath, Config)
	if err != nil {
		SugarLogger.Fatalf("加载配置文件失败: %s", err.Error())
	}
	SugarLogger.Debugf("读取到配置文件：%s", Config)
}

func GetAESKey() ([]byte, error) {
	return base64.URLEncoding.DecodeString(Config.EncodingAESKey + "==")
}
