package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type File struct {
	Name string `toml:"name"`
}

func LoadFile(filePath string) error {

	var code_id = "300607"
	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}

	return nil
}
