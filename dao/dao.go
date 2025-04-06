package dao

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacknotes/go-shell/conf"
)

var (
	// 定义对象是满足该接口的实例
	service *ServiceImpl
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

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{
		db: conf.C().Mysql.GetDB(),
	}
}

type ServiceImpl struct {
	db *sql.DB
}

func WriteDB(config *conf.Config) error {
	service = NewServiceImpl()
	// 1. 创建HTTP客户端
	client := &http.Client{}

	// 2. 构造请求参数
	url := "http://zxtp.guosen.com.cn:7615/TQLEX?Entry=CWServ.tdxf10_gg_jyds"

	for j := range config.App.Code {
		body := fmt.Sprintf("{\"Params\":[\"%s\",\"zjlx\",\"\"]}", config.App.Code[j])
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

		// 4. 处理响应
		var responseData ResponseData
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			return err
		}

		// 5. 写入MySQL
		tx, err := service.db.Begin()
		if err != nil {
			return err
		}
		stmt, err := tx.Prepare(InsertSQL)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, data := range responseData.ResultSets {
			for i := range data.Content {
				// fmt.Printf("debug %d,%s,%s,%s,%s,%s,%s,%s,%s,%s", file.App.Code[j], data.Content[i][0], data.Content[i][1], data.Content[i][2], data.Content[i][3], data.Content[i][4], data.Content[i][5], data.Content[i][6], data.Content[i][7], data.Content[i][8])
				_, err := stmt.Exec(config.App.Code[j], data.Content[i][0], data.Content[i][1], data.Content[i][2], data.Content[i][3], data.Content[i][4], data.Content[i][5], data.Content[i][6], data.Content[i][7], data.Content[i][8])
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
	}

	return nil
}

func NewDefaultData() *Data {
	return &Data{}
}

type Data struct {
	Code  string
	Date  string
	JLR   string
	ZLJLR string
}

func SelectData(config *conf.Config) error {
	// 创建文件写入对象（追加模式）
	// file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// 创建文件写入对象（覆盖模式）
	file, err := os.OpenFile(config.App.OutFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 确保缓冲区数据写入磁盘

	queryStmt, err := service.db.Prepare(SelectSQL)
	if err != nil {
		return err
	}
	defer queryStmt.Close()

	// QuanTianZhuMaiJinE
	rows, err := queryStmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	writer.WriteString("# HELP QuanTianZhuMaiJinE value" + "\n" + "# TYPE QuanTianZhuMaiJinE gauge" + "\n")
	for rows.Next() {
		ins := NewDefaultData()
		err := rows.Scan(
			&ins.Code, &ins.Date, &ins.JLR, &ins.ZLJLR,
		)
		if err != nil {
			return err
		}
		// prometheus格式，key不能包含'"',k v之前使用=号隔开，不能使用':'号隔开
		formatString := fmt.Sprintf("QuanTianZhuMaiJinE{Code=\"%s\",Date=\"%s\",JLR=\"%s\",ZLJLR=\"%s\"} %s", ins.Code, ins.Date, ins.JLR, ins.ZLJLR, ins.JLR)
		_, writeErr := writer.WriteString(formatString + "\n")
		if writeErr != nil {
			return writeErr
		}
	}

	// QuanTianZhuLiJinE
	rows, err = queryStmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	writer.WriteString("\n" + "# HELP QuanTianZhuLiJinE value" + "\n" + "# TYPE QuanTianZhuLiJinE gauge" + "\n")
	for rows.Next() {
		ins := NewDefaultData()
		err := rows.Scan(
			&ins.Code, &ins.Date, &ins.JLR, &ins.ZLJLR,
		)
		if err != nil {
			return err
		}

		// prometheus格式，key不能包含'"',k v之前使用=号隔开，不能使用':'号隔开
		formatString := fmt.Sprintf("QuanTianZhuLiJinE{Code=\"%s\",Date=\"%s\",JLR=\"%s\",ZLJLR=\"%s\"} %s", ins.Code, ins.Date, ins.JLR, ins.ZLJLR, ins.ZLJLR)
		_, writeErr := writer.WriteString(formatString + "\n")
		if writeErr != nil {
			return writeErr
		}
	}

	return nil
}
