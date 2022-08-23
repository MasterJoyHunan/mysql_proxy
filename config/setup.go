package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	LogConf        = Log{}
	ProxyMysqlConf = ProxyMysql{}
	RealMysqlConf  = RealMysql{}
)

func Setup(fileName string) {
	viper.SetConfigFile(fileName)

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("读取配置文件错误", err)
	}

	if err := viper.UnmarshalKey("Log", &LogConf); err != nil {
		log.Panic("Log配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("ProxyMysql", &ProxyMysqlConf); err != nil {
		log.Panic("ProxyMysql 配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("RealMysql", &RealMysqlConf); err != nil {
		log.Panic("RealMysql 配置文件格式错误", err)
	}
}
