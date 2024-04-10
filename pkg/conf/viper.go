package conf

import (
	"billing_finops/pkg/sdk"
	"billing_finops/utils"
	"github.com/spf13/viper"
)

func Init(cfg string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		utils.Log.Panic("config file path is empty")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			utils.Log.Panic("config file not found")
		} else {
			utils.Log.Panic("config file read error")
		}
		return
	}

	// 获得 运营商 列表
	lists := viper.Get("clouds.list")
	for _, v := range lists.([]interface{}) {
		item := v.(map[string]interface{})
		name := item["name"].(string)
		auth := item["auth"].(map[string]interface{})
		accessKeyId := auth["access_key_id"].(string)
		accessKeySecret := auth["access_key_secret"].(string)
		utils.Log.Infof("read %s config success", name)
		// 初始化运营商客户端sdk
		initSdk(name, accessKeyId, accessKeySecret)
	}

}

func initSdk(name string, id string, secret string) {
	if name == "aliyun" {
		// 初始化阿里云sdk
		err := sdk.NewClient(id, secret)
		if err != nil {
			utils.Log.Panic("can not create aliyun client:", err)
		}
	}
	// todo 其他运营商

}
