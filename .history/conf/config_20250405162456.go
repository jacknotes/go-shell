package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var GetFile *File

func NewDefaultFile() *File {
	return &File{
		Code: []int{300607},
	}
}

type File struct {
	Code []int `toml:"Code"`
}

func LoadConfigFromToml(filePath string) error {
	GetFile = NewDefaultFile()

	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, GetFile)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}
	fmt.Printf("debug %d", GetFile.Code[0])
	return nil
}
