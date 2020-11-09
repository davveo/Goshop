package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"log"
)

func CreateWechatMessageTemplateFactory(sqlType string) *WechatMessageTemplateModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &WechatMessageTemplateModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("CustomWordsModel工厂初始化失败")
	return nil
}

type WechatMessageTemplateModel struct {
	*BaseModel
}

func (wmtm *WechatMessageTemplateModel) list(pageNo, pageSize int) []map[string]interface{} {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_wechat_msg_template  ")

	if pageNo != 0 && pageSize != 0 {
		sqlString.WriteString(" order by count desc")
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := wmtm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}

	return tableData
}

func (wmtm *WechatMessageTemplateModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	pageNo, _ := params["page_no"].(int)
	pageSize, _ := params["page_size"].(int)
	return wmtm.list(pageNo, pageSize), wmtm.count()
}

func (wmtm *WechatMessageTemplateModel) count() (rows int64) {
	var (
		sql = "select count(1) from es_wechat_msg_template"
	)

	err := wmtm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (wmtm *WechatMessageTemplateModel) IsSync() bool {
	if len(wmtm.list(0, 0)) == 0 {
		return true
	}
	return false
}

func (wmtm *WechatMessageTemplateModel) Sync() bool {
	return false
}

func (wmtm *WechatMessageTemplateModel) getModel(Id int) (map[string]interface{}, error) {
	sqlString := "select * from es_wechat_msg_template where id = ?"
	rows := wmtm.QuerySql(sqlString, Id)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, err
	}
	var tmp map[string]interface{}
	if len(tableData) > 0 {
		tmp = tableData[0]
	}
	return tmp, nil
}
