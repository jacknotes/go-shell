package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)



func NewDefaultFile() *App {
	return &File{
		Name: "300607"
	}
}

type File struct {
	Name string `toml:"name"`
}

func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()

	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}

	return nil
}
