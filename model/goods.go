package model

import (
	"bytes"
	"errors"
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
	GoodsId             int    `json:"goods_id"`              // 主键
	GoodsName           string `json:"goods_name"`            // 商品名称
	Sn                  string `json:"sn"`                    // 商品编号
	BrandId             int    `json:"brand_id"`              // 品牌id
	CategoryId          int    `json:"category_id"`           // 分类id
	GoodsType           string `json:"goods_type"`            // 商品类型normal普通point积分
	Weight              string `json:"weight"`                // 重量
	MarketEnable        int    `json:"market_enable"`         // 上架状态 1上架  0下架
	Intro               string `json:"intro"`                 // 详情
	Price               string `json:"price"`                 // 商品价格
	Cost                string `json:"cost"`                  // 成本价格
	Mktprice            string `json:"mktprice"`              // 市场价格
	HaveSpec            int    `json:"have_spec"`             // 是否有规格0没有 1有
	CreateTime          int64  `json:"create_time"`           // 创建时间
	LastModify          int64  `json:"last_modify"`           // 最后修改时间
	ViewCount           int    `json:"view_count"`            // 浏览数量
	BuyCount            int    `json:"buy_count"`             // 购买数量
	Disabled            int    `json:"disabled"`              // 是否被删除0 删除 1未删除
	Quantity            int    `json:"quantity"`              // 库存
	EnableQuantity      int    ` json:"enable_quantity"`      // 可用库存
	Point               int    `json:"point"`                 // 如果是积分商品需要使用的积分
	PageTitle           string `json:"page_title"`            // seo标题
	MetaKeywords        string `json:"meta_keywords"`         // seo关键字
	MetaDescription     string `json:"meta_description"`      // seo描述
	Grade               string `json:"grade"`                 // 商品好评率
	Thumbnail           string ` json:"thumbnail"`            // 缩略图路径
	Big                 string `json:"big"`                   // 大图路径
	Small               string `json:"small"`                 // 小图路径
	Original            string ` json:"original"`             // 原图路径
	SellerId            int    ` json:"seller_id"`            // 卖家id
	ShopCatId           int    `json:"shop_cat_id"`           // 店铺分类id
	CommentNum          int    `json:"comment_num"`           // 评论数量
	TemplateId          int    `json:"template_id"`           // 运费模板id
	GoodsTransfeeCharge int    `json:"goods_transfee_charge"` // 谁承担运费0：买家承担，1：卖家承担
	SellerName          string `json:"seller_name"`           // 卖家名字
	IsAuth              int    `json:"is_auth"`               // 0 需要审核 并且待审核，1 不需要审核 2需要审核 且审核通过 3 需要审核 且审核未通过
	AuthMessage         string `json:"auth_message"`          // 审核信息
	SelfOperated        int    `json:"self_operated"`         // 是否是自营商品 0 不是 1是
	UnderMessage        string `json:"under_message"`         // 下架原因
	Priority            int    `json:"priority"`              // 优先级:高(3)、中(2)、低(1)
	CategoryName        string `json:"category_name"`         // 优先级:高(3)、中(2)、低(1)
}

type GoodsModel struct {
	*BaseModel
}

func (gm *GoodsModel) NewGoods(length int) (allGoodsList []Goods) {
	var (
		sqlString = "select * from es_goods where market_enable = 1 and " +
			"disabled = 1 order by create_time desc limit 0, ?"
	)

	rows := gm.QuerySql(sqlString, length)
	defer rows.Close()

	if rows != nil {
		for rows.Next() {
			goods := Goods{}
			err := sql_utils.ParseToStruct(rows, &goods)
			if err != nil {
				log.Println(err)
			}
			allGoodsList = append(allGoodsList, goods)
		}
		_ = rows.Close()
	}
	return allGoodsList
}

func (gm *GoodsModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var sqlString bytes.Buffer
	sqlString.WriteString("select g.goods_id,g.goods_name,g.sn,g.brand_id,g.thumbnail,g.seller_name," +
		"g.enable_quantity,g.quantity,g.price,g.create_time,g.market_enable,g.is_auth,g.under_message," +
		"g.priority from es_goods g where 1 = 1")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	gm.baseQuery(params, &sqlString)
	gm.categoryQuery(params, &sqlString)
	gm.shopCatQuery(params, &sqlString)

	sqlString.WriteString(" order by g.priority desc,g.create_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
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
	return tableData, gm.count()
}

func (gm *GoodsModel) baseQuery(params map[string]interface{}, sqlString *bytes.Buffer) {
	sn := params["sn"].(string)
	keyword := params["keyword"].(string)
	brandId := params["brand_id"].(string)
	disabled := params["disabled"].(string)
	endPrice := params["end_price"].(string)
	sellerId := params["seller_id"].(string)
	goodsName := params["goods_name"].(string)
	goodsType := params["goods_type"].(string)
	IsAuth, okIsAuth := params["is_auth"].(int)
	startPrice := params["start_price"].(string)
	sellerName := params["seller_name"].(string)
	marketEnable := params["market_enable"].(string)
	if okIsAuth {
		sqlString.WriteString(" and is_auth = ")
		sqlString.WriteString(strconv.Itoa(IsAuth))
	}
	if disabled == "" {
		disabled = "1"
	}
	if disabled != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.disabled = %s", disabled))
	}
	// 上下架
	if marketEnable != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.market_enable =  %s", marketEnable))
	}
	if keyword != "" {
		sqlString.WriteString(fmt.Sprintf(" and (g.goods_name like '%s' or g.sn like '%s' ) ", keyword, keyword))
	}
	if goodsName != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.goods_name like '%s'", goodsName))
	}
	if sellerName != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.seller_name like '%s'", sellerName))
	}
	if sn != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.sn like '%s'", sn))
	}
	if sellerId != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.seller_id = %s", sellerId))
	}
	if goodsType != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.goods_type = %s", goodsType))
	}
	if brandId != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.brand_id = %s", brandId))
	}
	if startPrice != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.price >= %s", startPrice))
	}
	if endPrice != "" {
		sqlString.WriteString(fmt.Sprintf(" and g.price <= %s", endPrice))
	}
}

func (gm *GoodsModel) categoryQuery(params map[string]interface{}, sqlString *bytes.Buffer) error {
	// 商城分类，同时需要查询出子分类的商品
	categoryPath := params["category_path"].(string)
	if categoryPath != "" {
		sql := fmt.Sprintf("select category_id from es_category where category_path like '%s'", categoryPath)
		rows := gm.QuerySql(sql)
		defer rows.Close()

		tableData, err := sql_utils.ParseJSON(rows)
		if err != nil {
			log.Println("sql_utils.ParseJSON 错误", err.Error())
			return err
		}
		if len(tableData) == 0 {
			return errors.New("分类不存在")
		}
		var tmp bytes.Buffer
		for index, data := range tableData {
			categoryId, ok := data["category_id"].(string)
			if index == 0 && ok {
				tmp.WriteString(fmt.Sprintf("%s", categoryId))
			} else {
				tmp.WriteString(fmt.Sprintf(",%s", categoryId))
			}
		}
		sqlString.WriteString(fmt.Sprintf(" and g.category_id in (%s)", tmp.String()))
	}
	return nil
}

func (gm *GoodsModel) shopCatQuery(params map[string]interface{}, sqlString *bytes.Buffer) {
	categoryPath := params["category_path"].(string)
	if categoryPath != "" {
		gm.getshopCatChidren(categoryPath)
	}
}

func (gm *GoodsModel) getshopCatChidren(categoryPath string) []string {

}

func (gm *GoodsModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_goods;"
	)

	err := gm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
