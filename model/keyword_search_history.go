package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateKeyWordSearchHistoryFactory(sqlType string) *KeyWordSearchHistoryModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &KeyWordSearchHistoryModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("CustomWordsModel工厂初始化失败")
	return nil
}

type KeyWordSearchHistoryModel struct {
	*BaseModel
}

func (kwshm *KeyWordSearchHistoryModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_keyword_search_history where count > 0 ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	keyWords, okKeywords := params["keywords"].(string)

	//按关键字查询
	if keyWords != "" && okKeywords {
		sqlString.WriteString(fmt.Sprintf(" and keyword like '%s'", "%"+keyWords+"%"))
	}

	sqlString.WriteString(" order by count desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := kwshm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, kwshm.count()
}

func (kwshm *KeyWordSearchHistoryModel) count() (rows int64) {
	var (
		sql = "select count(1) from es_keyword_search_history where count > 0"
	)

	err := kwshm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
