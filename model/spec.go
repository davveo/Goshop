package model

import (
	"bytes"
	"log"
	"Eshop/utils/sql_utils"
	"Eshop/utils/yml_config"
	"strconv"
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

func (s *SpecModel) List(params map[string]interface{}, keyword string) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_specification  where disabled = 1 and seller_id = 0 ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	if keyword != "" {
		sqlString.WriteString(" and spec_name like ")
		sqlString.WriteString(" '%" + keyword + "%' ")
	}

	sqlString.WriteString(" order by spec_id desc ")

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := s.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, s.count()
}

func (s *SpecModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_specification;"
	)

	err := s.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
