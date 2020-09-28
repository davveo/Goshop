package model

import (
	"bytes"
	"fmt"
	"log"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"strconv"
)

func CreateSiteNavigationFactory(sqlType string) *SiteNavigationModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &SiteNavigationModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type SiteNavigationModel struct {
	*BaseModel
}

func (sngm *SiteNavigationModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_site_navigation")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	clientType, okclientType := params["client_type"].(string)

	if clientType != "" && okclientType {
		sqlString.WriteString(fmt.Sprintf(" where client_type = '%s'", clientType))
	}
	sqlString.WriteString(" order by sort desc ")
	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := sngm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, sngm.count(clientType)
}

func (sngm *SiteNavigationModel) count(clientType string) (rows int64) {
	var (
		sql = fmt.Sprintf("select count(*) from es_site_navigation ")
	)

	if clientType != "" {
		sql += fmt.Sprintf(" where client_type = '%s'", clientType)
	}

	err := sngm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
