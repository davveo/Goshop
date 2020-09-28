package model

import (
	"fmt"
	"log"
	"Eshop/utils/sql_utils"
	"Eshop/utils/yml_config"
)

func CreateFocusPictureFactory(sqlType string) *FocusPictureModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &FocusPictureModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type FocusPictureModel struct {
	*BaseModel
}

func (fpm *FocusPictureModel) List(clientType string) ([]map[string]interface{}, error) {
	sqlString := fmt.Sprintf("select * from es_focus_picture  where client_type = '%s'order by id asc", clientType)
	rows := fpm.QuerySql(sqlString)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, err
	}

	return tableData, nil
}

func (fpm *FocusPictureModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_focus_picture"
	)

	err := fpm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
