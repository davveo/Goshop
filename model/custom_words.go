package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateCustomWordsFactory(sqlType string) *CustomWordsModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &CustomWordsModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("CustomWordsModel工厂初始化失败")
	return nil
}

type CustomWordsModel struct {
	*BaseModel
}

func (cwm *CustomWordsModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_custom_words where disabled = 1 ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	keyWords, okKeywords := params["keywords"].(string)

	//按关键字查询
	if keyWords != "" && okKeywords {
		sqlString.WriteString(fmt.Sprintf(" and name like '%s'", "%"+keyWords+"%"))
	}

	sqlString.WriteString(" order by modify_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := cwm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, cwm.count()
}

func (cwm *CustomWordsModel) count() (rows int64) {
	var (
		sql = "select count(1) from es_custom_words where disabled = 1"
	)

	err := cwm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
