package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"log"
	"strconv"
)

func CreateBillTotalFactory(sqlType string) *BillTotalModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &BillTotalModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("BillTotalModel工厂初始化失败")
	return nil
}

type BillTotalModel struct {
	*BaseModel
}

func (btm *BillTotalModel) List(query map[string]string) ([]map[string]interface{}, int64) {
	var sqlString bytes.Buffer

	pageNo, pageSize := query["page_no"], query["page_size"]
	sqlString.WriteString("select * from es_bill_total where 1=1 ")

	if pageNo != "" && pageSize != "" {
		pageNo, _ := strconv.Atoi(pageNo)
		pageSize, _ := strconv.Atoi(pageSize)
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	resSql := sqlString.String()
	rows := btm.QuerySql(resSql)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, btm.count(resSql)
}

func (btm *BillTotalModel) count(SqlString string) (rows int64) {
	err := btm.QueryRow(SqlString).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
