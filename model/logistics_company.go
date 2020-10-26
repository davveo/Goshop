package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateLogisticsCompanyFactory(sqlType string) *LogisticsCompanyModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &LogisticsCompanyModel{
			BaseModel: dbDriver,
		}
	}
	return nil
}

type LogisticsCompanyModel struct {
	*BaseModel
}

func (lcm *LogisticsCompanyModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_logistics_company")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	name, okName := params["name"].(string)
	status, okStatus := params["status"].(string)

	if status != "" && okStatus {
		sqlString.WriteString(fmt.Sprintf(" where delete_status = '%s'", status))
	}

	//按关键字查询
	if name != "" && okName {
		sqlString.WriteString(fmt.Sprintf(" and name like '%s'", "%"+name+"%"))
	}

	sqlString.WriteString(" order by id desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := lcm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, lcm.count()
}

func (lcm *LogisticsCompanyModel) count() (rows int64) {
	var (
		sql = "select count(1) from es_keyword_search_history where count > 0"
	)

	err := lcm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
