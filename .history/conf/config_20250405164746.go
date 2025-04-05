package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var GetFile *File

func NewDefaultFile() *File {
	return &File{
		App: NewDefaultApp(),
	}
}

type File struct {
	App   *App   `toml:"App"`
	Mysql *Mysql `toml:"Mysql"`
}

func NewDefaultApp() *App {
	return &App{
		Code: []int{300607},
	}
}

type App struct {
	Code []int `toml:"Code"`
}

func NewDefaultMysql() *Mysql {
	return &Mysql{
		Host:     "127.0.0.1",
		Port:     "3306",
		UserName: "test",
		Password: "123456",
		Database: "test",
	}
}

type Mysql struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
}

func LoadConfigFromToml(filePath string) error {
	GetFile = NewDefaultFile()

	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, GetFile)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}
	// fmt.Printf("debug %d", GetFile.App.Code[0])
	return nil
}
