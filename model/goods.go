package model

import (
	"bytes"
	"fmt"
	"log"
	"orange/utils/sql_utils"
	"orange/utils/yml_config"
	"strconv"
)

func CreateGoodsFactory(sqlType string) *GoodsModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &GoodsModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("goodsModel工厂初始化失败")
	return nil
}

type Goods struct {
	BrandID        string `sql:"brand_id" json:"brand_id"`
	CreateTime     string `sql:"create_time" json:"create_time"`
	EnableQuantity int64  `sql:"enable_quantity" json:"enable_quantity"`
	GoodsID        string `sql:"goods_id" json:"goods_id"`
	GoodsName      string `sql:"goods_name" json:"goods_name"`
	IsAuth         int  `sql:"is_auth" json:"is_auth"`
	MarketEnable   int  `sql:"market_enable" json:"market_enable"`
	Price          int  `sql:"price" json:"price"`
	Priority       int  `sql:"priority" json:"priority"`
	Quantity       int  `sql:"quantity" json:"quantity"`
	SellerName     string `sql:"seller_name" json:"seller_name"`
	Sn             string `sql:"sn" json:"sn"`
	Thumbnail      string `sql:"thumbnail" json:"thumbnail"`
	UnderMessage   string `sql:"under_message" json:"under_message"`
}

type GoodsModel struct {
	*BaseModel
}

func (gm *GoodsModel) NewGoods(length int) (allGoodsList []Goods) {

	sqlString := "select * from es_goods where market_enable = 1 and disabled = 1 order by create_time desc limit 0, ?"

	rows := gm.QuerySql(sqlString, length)
	defer rows.Close()

	if rows != nil {
		for rows.Next() {
			goods := Goods{}
			err := sql_utils.ParseToStruct(rows, &goods)
			fmt.Println(err)
			allGoodsList = append(allGoodsList, goods)
		}
		_ = rows.Close()
	}
	return allGoodsList
}

func (gm *GoodsModel) List(params map[string]interface{}) ([]map[string]interface{}, int) {
	var sqlString bytes.Buffer
	sqlString.WriteString("select g.goods_id,g.goods_name,g.sn,g.brand_id,g.thumbnail,g.seller_name," +
		"g.enable_quantity,g.quantity,g.price,g.create_time,g.market_enable,g.is_auth,g.under_message," +
		"g.priority from es_goods g where 1 = 1")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}
	rows := gm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}
	return tableData, 0
}
