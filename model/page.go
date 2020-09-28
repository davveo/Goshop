package model

import (
	"bytes"
	"fmt"
	"log"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"strconv"
)

func CreatePageFactory(sqlType string) *PageModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &PageModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type PageModel struct {
	*BaseModel
}

func (pm *PageModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_page ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := pm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, pm.count()
}

func (pm *PageModel) constructPageData(page *map[string]interface{}) {
	// 如果是WAP, 需要重新渲染
	// TODO
	// void constructPageData
}

func (pm *PageModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_page"
	)

	err := pm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (pm *PageModel) GetByType(clientType, pageType string) (map[string]interface{}, error) {
	sql := fmt.Sprintf("select * from es_page where client_type = '%s' and page_type = '%s'", clientType, pageType)

	rows := pm.QuerySql(sql)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, err
	}
	var tmp map[string]interface{}
	if len(tableData) >= 1 {
		tmp = tableData[0]
	}
	pm.constructPageData(&tmp)
	return tmp, nil
}
