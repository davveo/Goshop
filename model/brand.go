package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"log"
)

func CreateBrandFactory(sqlType string) *BrandModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &BrandModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("brandModel工厂初始化失败")
	return nil
}

type BrandModel struct {
	*BaseModel
	BrandId  int    `json:"brand_id"`
	Name     string `json:"name"`
	Disabled int    `json:"disabled"`
	Logo     string `json:"logo"`
}

func (bm *BrandModel) GetALllBrands() []map[string]interface{} {
	sqlString := "select * from es_brand order by brand_id desc "

	rows := bm.QuerySql(sqlString)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}
	return tableData
}

func (bm *BrandModel) GetCatBrand() {

}

func (bm *BrandModel) getModel(brandID int) *BrandModel {
	sql := "select brand_id, name, logo, disabled from es_brand where brand_id = ?"

	err := bm.QueryRow(sql, brandID).Scan(
		&bm.BrandId, &bm.Name, &bm.Logo, &bm.Disabled)

	if err == nil {
		return bm
	}
	return nil
}

func (bm *BrandModel) List(params map[string]interface{}, name string) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_brand ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	if name != "" {
		sqlString.WriteString(" where name like ")
		sqlString.WriteString(" '%" + name + "%' ")
	}

	sqlString.WriteString(" order by brand_id desc ")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := bm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, bm.count()
}

func (bm *BrandModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_brand;"
	)

	err := bm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (bm *BrandModel) Add() {

}

func (bm *BrandModel) Update() {

}

func (bm *BrandModel) Delete() {

}
