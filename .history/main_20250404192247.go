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
	dbTable = "zg_ag"
)

// 响应数据结构体（根据实际返回调整）
type ResponseData struct {
	ResultSets   []*ResultSet `json:"ResultSets"`
	ErrorCode    int64        `json:"ErrorCode"`
	ResultSetNum int64        `json:"ResultSetNum"`
}

type ResultSet struct {
	ColDes  []map[string]string `json:"ColDes"`
	ColNum  int64               `json:"ColNum"`
	Content [][]string          `json:"Content"`
	RowNum  int64               `json:"RowNum"`
}

// var code_id []string

// code_id=["300607"]

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
	stmt, err := tx.Prepare("INSERT INTO " + dbTable + " (code, date, jlr, jlrzb, zljlr, zljlrzb, cddjlr, cddjlrzb, ddjlr, ddjlrzb) VALUES (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, data := range responseData.ResultSets {
		for i, v := range data.Content {
			_, err := stmt.Exec(v[i])
			if err != nil {
				tx.Rollback()
				log.Fatal(err)
			}
		}
	}
	tx.Commit()
}
