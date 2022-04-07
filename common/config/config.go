package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var cfg = &Config{}

type Config struct {
	People     []string `json:"people"`       // 人员名单
	PerUserNum int32    `json:"per_user_num"` // 每次排班人数
	Num        int32    `json:"num"`          // 从第几个开始
}

func readConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, cfg)
}

func GetConfig() *Config {
	tmp := *cfg
	return &tmp
}
