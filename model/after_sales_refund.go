package model

import (
	"bytes"
	"log"
	"orange/utils/sql_utils"
	"orange/utils/yml_config"
	"strconv"
)

func CreateAfterSalesRefundFactory(sqlType string) *AfterSalesRefundModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSalesRefundModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type AfterSalesRefundModel struct {
	*BaseModel
}

func (asfm *AfterSalesRefundModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_refund where disabled = ? ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	sqlString.WriteString(" order by create_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := asfm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, asfm.count()
}

func (asfm *AfterSalesRefundModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_refund;"
	)

	err := asfm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
