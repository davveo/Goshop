package model

import (
	"bytes"
	"log"
	"orange/utils/sql_utils"
	"orange/utils/yml_config"
	"strconv"
)

func CreateGoodsFactory(sqlType string) *GoodsModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
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

type GoodsModel struct {
	*BaseModel
	BrandID        string `json:"brand_id"`
	CreateTime     string `json:"create_time"`
	EnableQuantity int64  `json:"enable_quantity"`
	GoodsID        string `json:"goods_id"`
	GoodsName      string `json:"goods_name"`
	IsAuth         int64  `json:"is_auth"`
	MarketEnable   int64  `json:"market_enable"`
	Price          int64  `json:"price"`
	Priority       int64  `json:"priority"`
	Quantity       int64  `json:"quantity"`
	SellerName     string `json:"seller_name"`
	Sn             string `json:"sn"`
	Thumbnail      string `json:"thumbnail"`
	UnderMessage   string `json:"under_message"`
}

func (u *GoodsModel) NewGoods(length int) {
	//sql := "select * from es_goods where market_enable = 1 and disabled = 1 order by create_time desc limit 0, ?"
	//rows := u.QuerySql(sql, length)
	//if rows != nil {
	//	tmp := make([]UsersModel, 0)
	//	for rows.Next() {
	//		err := rows.Scan(&u.Id, &u.UserName, &u.Department)
	//		if err == nil {
	//			tmp = append(tmp, *u)
	//		} else {
	//			log.Println("sql查询错误", err.Error())
	//		}
	//	}
	//	_ = rows.Close()
	//	return tmp
	//}
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
