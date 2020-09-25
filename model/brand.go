package model

import (
	"log"
	"orange/utils/sql_utils"
	"orange/utils/yml_config"
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

func (bm *BrandModel) List() {

}

func (bm *BrandModel) Add() {

}

func (bm *BrandModel) Update() {

}

func (bm *BrandModel) Delete() {

}
