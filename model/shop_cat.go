package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"fmt"
	"log"
)

func CreatGoshopCateGoryFactory(sqlType string) *ShopCateGoryModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
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

	sql := fmt.Sprintf("select shop_cat_id from es_shop_cat where cat_path like '%s' ", "%"+catPath+"%")
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
