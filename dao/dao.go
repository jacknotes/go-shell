package dao

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacknotes/go-shell/conf"
)

var (
	// 定义对象是满足该接口的实例
	service *ServiceImpl
	td      time.Duration = 300 * time.Millisecond
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
		fmt.Printf("debug jlr %s\n", config.App.Code[j])
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

		for _, data := range responseData.ResultSets[0].Content {
			// fmt.Printf("debug %s,%s,%s,%s,%s,%s,%s,%s,%s,%s", config.App.Code[j], data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8])
			_, err := stmt.Exec(config.App.Code[j], data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9])
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
		time.Sleep(td)
	}

	return nil
}

func WriteDB_SDGD(config *conf.Config) error {
	service = NewServiceImpl()
	// 1. 创建HTTP客户端
	client := &http.Client{}
	// 2. 构造请求参数
	url := "http://zxtp.guosen.com.cn:7615/TQLEX?Entry=CWServ.tdxf10_gg_gdyj"

	for j := range config.App.Code {
		fmt.Printf("debug sdgd %s\n", config.App.Code[j])
		body_sdgd := fmt.Sprintf("{\"Params\":[\"%s\",\"ltgd\",\"\",\"\",\"1\",\"1\",\"20\"]}", config.App.Code[j])
		jsonData_sdgd := []byte(body_sdgd)

		// 3. 发送POST请求
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData_sdgd))
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
		stmt, err := tx.Prepare(InsertSDGDSQL)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for i, data := range responseData.ResultSets {
			if i == 1 {
				// fmt.Println(data.Content[0])
				for k := range data.Content {
					// fmt.Println(data.Content[k])
					// fmt.Printf("debug %s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s", config.App.Code[j], data.Content[k][0], data.Content[k][1], data.Content[k][2], data.Content[k][3], data.Content[k][4], data.Content[k][5], data.Content[k][6], data.Content[k][7], data.Content[k][8], data.Content[k][9])
					_, err := stmt.Exec(config.App.Code[j], data.Content[k][0], data.Content[k][1], data.Content[k][2], data.Content[k][3], data.Content[k][4], data.Content[k][5], data.Content[k][6], data.Content[k][7], data.Content[k][8], data.Content[k][9])
					if err != nil {
						tx.Rollback()
						return err
					}
				}
			}
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
		time.Sleep(td)
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

	// QuanTianZhuMaiJinE
	str := "('"
	// in ('300041','300046')
	for i := range config.App.Code {
		if i == len(config.App.Code)-1 {
			str = str + config.App.Code[i] + "')"
		} else {
			str = str + config.App.Code[i] + "','"
		}

	}
	SelectSQL := fmt.Sprintf("SELECT code,date,jlr,zljlr FROM `zg_ag` WHERE code IN %s", str)
	// fmt.Printf("debug %s", SelectSQL)
	// return nil

	queryStmt, err := service.db.Prepare(SelectSQL)
	if err != nil {
		return err
	}
	defer queryStmt.Close()

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
