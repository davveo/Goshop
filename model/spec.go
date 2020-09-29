package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateSpecFactory(sqlType string) *SpecModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &SpecModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type SpecModel struct {
	*BaseModel
}

func (s *SpecModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString      bytes.Buffer
		countSqlString string
	)

	sqlString.WriteString("select * from es_specification  where disabled = 1 and seller_id = 0 ")

	keyword, okKeyword := params["keyword"].(string)
	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	if keyword != "" && okKeyword {
		sqlString.WriteString(" and spec_name like ")
		sqlString.WriteString(" '%" + keyword + "%' ")
	}

	sqlString.WriteString(" order by spec_id desc ")
	countSqlString = sql_utils.GetCountSql(sqlString.String())

	if okPageNo && okPageSize {
		sqlString.WriteString(fmt.Sprintf(" limit %d, %d", pageNo-1, pageSize))
	}

	rows := s.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, s.count(countSqlString)
}

func (s *SpecModel) count(sql string) (rows int64) {

	err := s.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
