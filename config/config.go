package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	People     []string `json:"people"`       // 人员名单
	PerUserNum int32    `json:"per_user_num"` // 每次排班人数
	Num        int32    `json:"num"`          // 从第几个开始
}

func NewConfig(path string) (cfg *Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	cfg = &Config{}
	err = json.Unmarshal(content, cfg)
	return
}

func (c *Config) Reload() {

}
