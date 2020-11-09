package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateMessageTemplateFactory(sqlType string) *MessageTemplateModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &MessageTemplateModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("CustomWordsModel工厂初始化失败")
	return nil
}

type MessageTemplateModel struct {
	*BaseModel
}

func (mtm *MessageTemplateModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_message_template")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	tType, okType := params["type"].(string)

	if tType != "" && okType {
		sqlString.WriteString(fmt.Sprintf(" where type = '%s'", tType))
	}

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := mtm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, mtm.count(tType)
}

func (mtm *MessageTemplateModel) count(tType string) (rows int64) {
	var (
		sql = "select count(1) from es_message_template where type = ? "
	)

	err := mtm.QueryRow(sql, tType).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
