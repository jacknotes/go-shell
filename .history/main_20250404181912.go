package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库配置
const (
	dbUser  = "go_shell"
	dbPass  = "go_shell"
	dbName  = "go_shell"
	dbHost  = "192.168.31.202:3306"
	dbTable = "go_shell"
)

// 响应数据结构体（根据实际返回调整）
type ResponseData struct {
	ResultSets []struct {
		Field1 string `json:"field1"`
		Field2 string `json:"field2"`
	} `json:"data"`,
	ErrorCode int,
	
}

func main() {
	// 1. 创建HTTP客户端
	client := &http.Client{}

	// 2. 构造请求参数
	url := "http://zxtp.guosen.com.cn:7615/TQLEX?Entry=CWServ.tdxf10_gg_jyds"
	jsonData := []byte(`{"Params":["300607","zjlx",""]}`)

	// 3. 发送POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 4. 处理响应
	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		log.Fatal(err)
	}

	// 5. 写入MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO " + dbTable + " (field1, field2) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, data := range responseData.Data {
		_, err := stmt.Exec(data.Field1, data.Field2)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}
	tx.Commit()
}
