package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	config *Config
	db     *sql.DB
)

func C() *Config {
	return config
}

func NewDefaultConfig() *Config {
	return &Config{
		Mysql: NewDefaultMysql(),
		App:   NewDefaultApp(),
	}
}

type Config struct {
	App   *App   `toml:"App"`
	Mysql *Mysql `toml:"Mysql"`
}

func NewDefaultApp() *App {
	return &App{
		Code:    []string{},
		OutFile: "output.txt",
	}
}

type App struct {
	Code    []string `toml:"Code"`
	OutFile string   `toml:"OutFile"`
}

func NewDefaultMysql() *Mysql {
	return &Mysql{
		Host:        "127.0.0.1:3306",
		UserName:    "test",
		Password:    "123456",
		Database:    "test",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

type Mysql struct {
	Host        string `toml:"host" env:"MYSQL_HOST"`
	UserName    string `toml:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" env:"MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
	lock        sync.Mutex
}

func (m *Mysql) GetDB() *sql.DB {
	// 直接加锁, 锁住临界区
	m.lock.Lock()
	defer m.lock.Unlock()

	// 如果实例不存在, 就初始化一个新的实例
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			panic(err)
		}
		db = conn
	}

	// 全局变量db就一定存在了
	return db
}

func (m *Mysql) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
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
