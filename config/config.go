package config

import (
	"log"
	"os"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Api_Key string
	Town_Id string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		//エラーの場合、時間とエラー内容を出力
		log.Printf("Failed to read file: %v", err)
		//logが読み込めない場合エラーコード1で終了
		os.Exit(1)
	}

	Config = ConfigList{
		Api_Key: cfg.Section("weatherapi").Key("api_key").String(),
		Town_Id: cfg.Section("weatherapi").Key("town_id").String(),
	}
}
