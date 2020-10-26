package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"errors"
	"fmt"
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

func (bm *BrandModel) GetModel(brandID int) (map[string]interface{}, error) {
	sqlString := "select brand_id, name, logo, disabled from es_brand where brand_id = ?"
	rows := bm.QuerySql(sqlString, brandID)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, err
	}
	var tmp map[string]interface{}
	if len(tableData) > 0 {
		tmp = tableData[0]
	}
	return tmp, nil
}

func (bm *BrandModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_brand ")

	name, okName := params["name"].(string)
	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	if name != "" && okName {
		sqlString.WriteString(fmt.Sprintf(" where name like '%s'", "%"+name+"%"))
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

func (bm *BrandModel) Add(params map[string]interface{}) (map[string]interface{}, error) {
	var (
		brandId int64
	)

	name := params["name"].(string)
	logo := params["logo"].(string)
	disabled := params["disabled"].(string)

	sqlString := "select * from es_brand where name = ? "
	rows := bm.QuerySql(sqlString, name)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, err
	}

	if len(tableData) != 0 {
		return nil, errors.New("品牌名称重复")
	}

	sqlString = "insert into `es_brand` (`name`, `logo`, `disabled`) values (?,?,?)"

	if brandId = bm.LastInsertId(sqlString, name, logo, disabled); brandId == -1 {
		return nil, errors.New("插入分类失败")
	}

	params["brand_id"] = brandId
	params["disabled"] = 1
	return params, nil
}

func (bm *BrandModel) Edit(params map[string]interface{}) (map[string]interface{}, error) {
	var (
		sqlString string
		name      = params["name"].(string)
		logo      = params["logo"].(string)
		brandId   = params["brand_id"].(int)
		disabled  = params["disabled"].(string)
	)

	brand, err := bm.GetModel(brandId)
	if brand == nil || err != nil {
		return nil, errors.New("品牌不存在")
	}
	sqlString = "select * from es_brand where name = ? and brand_id != ? "
	rows := bm.QuerySql(sqlString, name, brandId)
	defer rows.Close()

	brandList, _ := sql_utils.ParseJSON(rows)

	if len(brandList) > 0 {
		return nil, errors.New("品牌名称重复")
	}
	sqlString = "update  es_brand set `name` = ?, `logo` = ?, `disabled` = ? where brand_id = ?"
	if affected := bm.ExecuteSql(sqlString, name, logo, disabled, brandId); affected == -1 {
		return nil, errors.New("更新失败")
	}

	return brand, nil
}

func (bm *BrandModel) Delete(brandIds []int) error {
	var (
		err              error
		checkSql         string
		sqlLast          string
		sqlString        string
		countForCategory int64
		countForGoods    int64
	)
	idsStr := sql_utils.InSqlStr(brandIds)
	sqlString = "select count(0) from es_category_brand where brand_id in (" + idsStr + ")"

	err = bm.QueryRow(sqlString).Scan(&countForCategory)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}
	//检测是否有分类关联
	if countForCategory > 0 {
		return errors.New("已有分类关联，不能删除")
	}
	// 检测是否有商品关联
	checkSql = "select count(0) from es_goods where (disabled = 1 or disabled = 0) and brand_id in (" + idsStr + ")"
	err = bm.QueryRow(checkSql).Scan(&countForGoods)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}
	if countForGoods > 0 {
		return errors.New("已有商品关联，不能删除")
	}

	sqlLast = "delete from es_brand where brand_id in (" + idsStr + ") "
	if affected := bm.ExecuteSql(sqlLast); affected == -1 {
		return errors.New("删除标签失败")
	}
	return nil
}

func (bm *BrandModel) getBrandsByCategory(id int) []map[string]interface{} {
	sql := "select b.brand_id,b.`name`,b.logo " +
		"from es_category_brand cb inner join es_brand b on cb.brand_id=b.brand_id " +
		"where cb.category_id=? and b.disabled = 1 "

	rows := bm.QuerySql(sql, id)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}
	return tableData
}
