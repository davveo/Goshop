package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"log"
)

func CreateSpecFactory(sqlType string) *SpecModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
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

	pageNo, okPageNo := params["page_no"].(int)
	keyword, okKeyword := params["keyword"].(string)
	pageSize, okPageSize := params["page_size"].(int)

	if keyword != "" && okKeyword {
		sqlString.WriteString(sql_utils.Like("spec_name", keyword, true))
	}

	sqlString.WriteString(sql_utils.OrderBy("spec_id", "desc"))
	countSqlString = sql_utils.GetCountSql(sqlString.String())

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := s.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, sql_utils.Count(countSqlString, s.dbDriverRead)
}
