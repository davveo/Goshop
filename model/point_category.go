package model

import (
	"bytes"
	"fmt"
	"log"
	"orange/utils/sql_utils"
	"orange/utils/yml_config"
)

func CreatePointCateGoryFactory(sqlType string) *PointCateGory {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &PointCateGory{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type PointCateGory struct {
	*BaseModel
}

func (pcg *PointCateGory) List(params map[string]interface{}) []map[string]interface{} {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_exchange_cat")

	parentID := params["parent_id"].(string)

	if parentID != "" {
		sqlString.WriteString(fmt.Sprintf(" where parent_id = %s", parentID))
	}
	sqlString.WriteString(" order by category_order asc ")

	rows := pcg.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}

	return tableData
}

func (pcg *PointCateGory) count() (rows int64) {
	var (
		sql = "select count(*) from es_receipt_history;"
	)

	err := pcg.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
