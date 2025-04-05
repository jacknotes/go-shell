package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func NewDefaultFile() *File {
	return &File{
		Code: []int{300607},
	}
}

type File struct {
	Code []int `toml:"Code"`
}

func LoadConfigFromToml(filePath string) error {
	config = NewDefaultFile()

	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}

	return nil
}
