package dao

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacknotes/go-shell.git/conf"
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

func WriteDB(file *conf.File) error {

	// 1. 创建HTTP客户端
	client := &http.Client{}

	// 2. 构造请求参数
	url := "http://zxtp.guosen.com.cn:7615/TQLEX?Entry=CWServ.tdxf10_gg_jyds"

	for i := range file.App.Code {
		body := fmt.Sprintf("{\"Params\":[\"%d\",\"zjlx\",\"\"]}", file.App.Code[i])
		jsonData := []byte(body)

		// 3. 发送POST请求
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

	}

	// 4. 处理响应
	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return err
	}

	// 5. 写入MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", file.Mysql.UserName,
		file.Mysql.Password, file.Mysql.Host, file.Mysql.Database))
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO zg_ag" + " (code, date, jlr, jlrzb, zljlr, zljlrzb, cddjlr, cddjlrzb, ddjlr, ddjlrzb) VALUES (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, data := range responseData.ResultSets {
		for i := range data.Content {
			_, err := stmt.Exec(code_id, data.Content[i][0], data.Content[i][1], data.Content[i][2], data.Content[i][3], data.Content[i][4], data.Content[i][5], data.Content[i][6], data.Content[i][7], data.Content[i][8])
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
