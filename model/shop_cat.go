package model

import (
	"fmt"
	"log"
	"orange/utils/sql_utils"
	"orange/utils/yml_config"
)

func CreateShopCateGoryFactory(sqlType string) *ShopCateGoryModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &ShopCateGoryModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type ShopCateGoryModel struct {
	*BaseModel
}

func (scgm *ShopCateGoryModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	return nil, scgm.count()
}

func (scgm *ShopCateGoryModel) getChildren(catPath string) ([]map[string]interface{}, error) {

	sql := fmt.Sprintf(""+
		"select shop_cat_id from es_shop_cat where cat_path like '%s' ", catPath)
	rows := scgm.QuerySql(sql)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, err
	}
	return tableData, nil
}

func (scgm *ShopCateGoryModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_shop_cat;"
	)

	err := scgm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
