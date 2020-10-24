package model

import (
	"Goshop/global/consts"
	"Goshop/utils/rabbitmq"
	"Goshop/utils/sql_utils"
	"Goshop/utils/time_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateGoodsFactory(ctx *gin.Context, sqlType string) *GoodsModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	mq := rabbitmq.GetRabbitmq()
	if mq == nil {
		log.Fatal("goodsModel mq初始化失败")
	}
	amqpTemplate, err := mq.Producer("goods")
	if err != nil {
		log.Fatal("goodsModel producer初始化失败")
	}

	if dbDriver != nil {
		return &GoodsModel{
			BaseModel:    dbDriver,
			ctx:          ctx,
			amqpTemplate: amqpTemplate,
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
	ctx          *gin.Context
	amqpTemplate *rabbitmq.Producer
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
			_ = sql_utils.ParseToStruct(rows, &goods)
			allGoodsList = append(allGoodsList, goods)
		}
		_ = rows.Close()
	}
	return allGoodsList
}

func (gm *GoodsModel) Up(goodsId int) error {
	sqlString := "select disabled,market_enable,seller_id from es_goods where goods_id = ?"
	rows := gm.QuerySql(sqlString, goodsId)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return err
	}
	var tmp map[string]interface{}
	if len(tableData) > 0 {
		tmp = tableData[0]
	}
	disabled := tmp["disabled"].(int64)
	marketEnable := tmp["market_enable"].(int64)
	sellerId := tmp["seller_id"].(int64)
	operateAllowable := NewOperateAllowable(marketEnable, disabled)

	//查询店铺是否是关闭中，若未开启，则不能上架
	shop, _ := CreateShopFactory("").GetShop(sellerId)
	if shop == nil || shop["shop_disable"] != "OPEN" {
		return errors.New("店铺关闭中,商品不能上架操作")
	}
	//下架未删除才能上架
	if !operateAllowable.getAllowMarket() {
		return errors.New("商品不能上架操作")
	}

	sqlString = "update es_goods set market_enable = 1 and disabled = 1 where goods_id  = ?"
	if gm.ExecuteSql(sqlString, goodsId) == -1 {
		return errors.New("上架商品更新失败")
	}
	rds.Remove(fmt.Sprintf("%s_%d", consts.GOODS, goodsId))
	// TODO
	/* 后面在完善逻辑
	GoodsChangeMsg goodsChangeMsg = new GoodsChangeMsg(new Integer[]{goodsId}, GoodsChangeMsg.UPDATE_OPERATION);
	this.amqpTemplate.convertAndSend(AmqpExchange.GOODS_CHANGE, AmqpExchange.GOODS_CHANGE + "_ROUTING", goodsChangeMsg);
	*/
	return nil
}

// TODO ctx优化
func (gm *GoodsModel) Under(goodsIds []int, reason string, permission int) error {
	if len(reason) > 500 {
		return errors.New("下架原因长度不能超过500个字符")
	}
	idStr := sql_utils.InSqlStr(goodsIds)

	if permission == consts.PermissionSELLER {
		gm.checkPermission(goodsIds, consts.GoodsOperateUNDER)
		sellerUserName := gm.ctx.GetString("user_name")
		reason = "店员" + sellerUserName + "下架，原因为：" + reason
	} else {
		//查看是否是不能下架的状态
		sqlString := "select disabled,market_enable from es_goods where goods_id in (" + idStr + ")"
		rows := gm.QuerySql(sqlString)
		defer rows.Close()

		tableData, err := sql_utils.ParseJSON(rows)
		if err != nil {
			log.Println("sql_utils.ParseJSON 错误", err.Error())
			return err
		}
		for _, data := range tableData {
			disabled := data["disabled"].(int64)
			marketEnable := data["market_enable"].(int64)
			operateAllowable := NewOperateAllowable(marketEnable, disabled)

			//上架并且没有删除的可以下架
			if !operateAllowable.getAllowUnder() {
				return errors.New("存在不能下架的商品，不能操作")
			}
		}
		reason = "平台下架，原因为：" + reason
	}
	sqlString := "update es_goods set market_enable = 0,under_message = ?, last_modify=?  where goods_id in (" + idStr + ")"

	if gm.ExecuteSql(sqlString, reason, time.Now().Unix()) == -1 {
		return errors.New("下架商品更新失败")
	}

	//清除相关的关联
	for _, goodsId := range goodsIds {
		gm.cleanGoodsAssociated(goodsId, 0)
	}

	/*TODO
	GoodsChangeMsg goodsChangeMsg = new GoodsChangeMsg(goodsIds, GoodsChangeMsg.UNDER_OPERATION, reason);
	this.amqpTemplate.convertAndSend(AmqpExchange.GOODS_CHANGE, AmqpExchange.GOODS_CHANGE + "_ROUTING", goodsChangeMsg);
	*/
	return nil
}

// 在商品删除、下架要进行调用
func (gm *GoodsModel) cleanGoodsAssociated(goodsId int, markEnable int) {
	if yml_config.CreateYamlFactory().GetBool("AppDebug") {
		log.Println("清除goodsid[" + string(goodsId) + "]相关的缓存，包括促销的缓存")
	}
	rds.Remove(fmt.Sprintf("%s_%d", consts.GOODS, goodsId))

	// 删除这个商品的sku缓存(必须要在删除库中sku前先删缓存),首先查出商品对应的sku_id
	sqlString := "select sku_id from es_goods_sku where goods_id = ?"
	rows := gm.QuerySql(sqlString, goodsId)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
	}
	for _, data := range tableData {
		skuId := data["sku_id"].(int64)
		rds.Remove(fmt.Sprintf("%s_%d", consts.SKU, skuId))
	}

	//不再读一次缓存竟然清不掉？？所以在这里又读了一下
	rds.Gain(fmt.Sprintf("%s_%d", consts.GOODS, goodsId))

	//删除该商品关联的活动缓存
	currTimeStr := time_utils.GetDateStr(consts.TimeFormatStyleV2)

	//清除此商品的缓存
	rds.Remove(fmt.Sprintf("%s_%s_%d", consts.PROMOTION_KEY, currTimeStr, goodsId))

	if markEnable == 0 {
		gm.deleteExchange(goodsId)
	}
}

// 删除积分商品
func (gm *GoodsModel) deleteExchange(goodsId int) {
	CreateExchangeFactory("").delete(goodsId)
}

// 查看商品是否属于当前登录用户
func (gm *GoodsModel) checkPermission(goodsIds []int, goodsOperate int) {
	sellerId := gm.ctx.GetString("user_id")
	idStr := sql_utils.InSqlStr(goodsIds)
	sqlString := "select disabled,market_enable from es_goods where goods_id in (" + idStr + ") and seller_id = ?"
	rows := gm.QuerySql(sqlString, sellerId)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
	}

	if len(tableData) != len(goodsIds) {
		log.Println("存在不属于您的商品，不能操作")
	}

	for _, data := range tableData {
		disabled := data["disabled"].(int64)
		marketEnable := data["market_enable"].(int64)
		operateAllowable := NewOperateAllowable(marketEnable, disabled)

		switch goodsOperate {
		case consts.GoodsOperateDELETE:
			//下架的删除了的才能还原
			if !operateAllowable.getAllowDelete() {
				log.Println("存在不能删除的商品，不能操作")
			}
			break
		case consts.GoodsOperateRECYCLE:
			//下架的商品才能放入回收站
			if !operateAllowable.getAllowRecycle() {
				log.Println("存在不能放入回收站的商品，不能操作")
			}
			break
		case consts.GoodsOperateREVRET:
			//下架的删除了的才能还原
			if !operateAllowable.getAllowRevert() {
				log.Println("存在不能还原的商品，不能操作")
			}
			break
		case consts.GoodsOperateUNDER:
			//上架并且没有删除的可以下架
			if !operateAllowable.getAllowUnder() {
				log.Println("存在不能下架的商品，不能操作")
			}
			break
		default:
			break
		}
	}
}

func (gm *GoodsModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		countSqlString string
		sqlString      bytes.Buffer
		err            error
	)
	sqlString.WriteString("select g.goods_id,g.goods_name,g.sn,g.brand_id,g.thumbnail,g.seller_name," +
		"g.enable_quantity,g.quantity,g.price,g.create_time,g.market_enable,g.is_auth,g.under_message," +
		"g.priority from es_goods g where 1 = 1")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	_ = gm.baseQuery(params, &sqlString)
	err = gm.categoryQuery(params, &sqlString)
	if err != nil {
		log.Println("gm.categoryQuery 错误", err.Error())
	}
	err = gm.shopCatQuery(params, &sqlString)
	if err != nil {
		log.Println("gm.shopCatQuery 错误", err.Error())
	}

	sqlString.WriteString(" order by g.priority desc,g.create_time desc")
	countSqlString = sql_utils.GetCountSql(sqlString.String())

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}
	rows := gm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}
	return tableData, gm.count(countSqlString)
}

func (gm *GoodsModel) baseQuery(params map[string]interface{}, sqlString *bytes.Buffer) error {
	sn, okSn := params["sn"].(string)
	keyword, okKeyword := params["keyword"].(string)
	brandId, okBrandID := params["brand_id"].(string)
	disabled, okDisabled := params["disabled"].(string)
	endPrice, okEndPrice := params["end_price"].(string)
	sellerId, okSellerId := params["seller_id"].(string)
	goodsName, okGoodsName := params["goods_name"].(string)
	goodsType, okGoodsType := params["goods_type"].(string)
	IsAuth, okIsAuth := params["is_auth"].(int)
	startPrice, okStartPrice := params["start_price"].(string)
	sellerName, okSellerName := params["seller_name"].(string)
	marketEnable, okMarketEnable := params["market_enable"].(string)
	if okIsAuth {
		sqlString.WriteString(" and is_auth = ")
		sqlString.WriteString(strconv.Itoa(IsAuth))
	}
	if disabled == "" {
		disabled = "1"
	}
	if disabled != "" && okDisabled {
		sqlString.WriteString(fmt.Sprintf(" and g.disabled = %s", disabled))
	}
	// 上下架
	if marketEnable != "" && okMarketEnable {
		sqlString.WriteString(fmt.Sprintf(" and g.market_enable =  %s", marketEnable))
	}
	if keyword != "" && okKeyword {
		sqlString.WriteString(fmt.Sprintf(" and (g.goods_name like '%s' or g.sn like '%s' ) ", "%"+keyword+"%", "%"+keyword+"%"))
	}
	if goodsName != "" && okGoodsName {
		sqlString.WriteString(fmt.Sprintf(" and g.goods_name like '%s'", "%"+goodsName+"%"))
	}
	if sellerName != "" && okSellerName {
		sqlString.WriteString(fmt.Sprintf(" and g.seller_name like '%s'", "%"+sellerName+"%"))
	}
	if sn != "" && okSn {
		sqlString.WriteString(fmt.Sprintf(" and g.sn like '%s'", "%"+sn+"%"))
	}
	if sellerId != "" && okSellerId {
		sqlString.WriteString(fmt.Sprintf(" and g.seller_id = %s", sellerId))
	}
	if goodsType != "" && okGoodsType {
		sqlString.WriteString(fmt.Sprintf(" and g.goods_type = %s", goodsType))
	}
	if brandId != "" && okBrandID {
		sqlString.WriteString(fmt.Sprintf(" and g.brand_id = %s", brandId))
	}
	if startPrice != "" && okStartPrice {
		sqlString.WriteString(fmt.Sprintf(" and g.price >= %s", startPrice))
	}
	if endPrice != "" && okEndPrice {
		sqlString.WriteString(fmt.Sprintf(" and g.price <= %s", endPrice))
	}
	return nil
}

func (gm *GoodsModel) categoryQuery(params map[string]interface{}, sqlString *bytes.Buffer) error {
	// 商城分类，同时需要查询出子分类的商品
	categoryPath, okCategoryPath := params["category_path"].(string)
	if categoryPath != "" && okCategoryPath {
		sql := fmt.Sprintf("select category_id from es_category where category_path like '%s'", "%"+categoryPath+"%")
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

func (gm *GoodsModel) shopCatQuery(params map[string]interface{}, sqlString *bytes.Buffer) error {
	categoryPath, okCategoryPath := params["category_path"].(string)
	if categoryPath != "" && okCategoryPath {
		catList, err := gm.getShopCatChidren(categoryPath)

		if err != nil {
			return err
		}

		if len(catList) <= 0 {
			return errors.New("店铺分组不存在")
		}
		var tmp bytes.Buffer
		for index, data := range catList {
			categoryId, ok := data["shop_cat_id"].(string)
			if index == 0 && ok {
				tmp.WriteString(fmt.Sprintf("%s", categoryId))
			} else {
				tmp.WriteString(fmt.Sprintf(",%s", categoryId))
			}
		}
		sqlString.WriteString(fmt.Sprintf(" and g.shop_cat_id in (%s)", tmp.String()))
	}
	return nil
}

func (gm *GoodsModel) getShopCatChidren(categoryPath string) ([]map[string]interface{}, error) {
	return CreatGoshopCateGoryFactory("").getChildren(categoryPath)
}

func (gm *GoodsModel) count(sql string) (rows int64) {

	err := gm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (gm *GoodsModel) Delete(goodsIds []int) error {
	gm.checkPermission(goodsIds, consts.GoodsOperateDELETE)
	idStr := sql_utils.InSqlStr(goodsIds)
	sqlString := "update es_goods set disabled = -1  where goods_id in (" + idStr + ")"
	if gm.ExecuteSql(sqlString) == -1 {
		return errors.New("删除商品失败")
	}
	goodsChangeMsg := rabbitmq.BuildMsg(map[string]interface{}{
		"message":        "",
		"goods_ids":      goodsIds,
		"operation_type": consts.OperationDelOperation,
	})
	err := gm.amqpTemplate.Publish(consts.ExchangeGoodsChange,
		consts.ExchangeGoodsChange+"_ROUTING", goodsChangeMsg)
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}
	return nil
}

func (gm *GoodsModel) Revert(goodsIds []int) error {
	gm.checkPermission(goodsIds, consts.GoodsOperateREVRET)
	idStr := sql_utils.InSqlStr(goodsIds)
	sqlString := "update  es_goods set disabled = 1  where goods_id in (" + idStr + ")"
	if gm.ExecuteSql(sqlString) == -1 {
		return errors.New("删除商品失败")
	}
	goodsChangeMsg := rabbitmq.BuildMsg(map[string]interface{}{
		"message":        "",
		"goods_ids":      goodsIds,
		"operation_type": consts.OperationRevertOperation,
	})
	err := gm.amqpTemplate.Publish(consts.ExchangeGoodsChange,
		consts.ExchangeGoodsChange+"_ROUTING", goodsChangeMsg)
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}
	return nil
}

func (gm *GoodsModel) InRecycle(goodsIds []int) error {
	gm.checkPermission(goodsIds, consts.GoodsOperateRECYCLE)
	idStr := sql_utils.InSqlStr(goodsIds)
	sqlString := "update  es_goods set disabled = 0 ,market_enable=0 , last_modify=?  where goods_id in (" + idStr + ")"
	if gm.ExecuteSql(sqlString) == -1 {
		return errors.New("删除商品失败")
	}

	//清除相关的关联
	for _, goodsId := range goodsIds {
		gm.cleanGoodsAssociated(goodsId, 0)
	}

	goodsChangeMsg := rabbitmq.BuildMsg(map[string]interface{}{
		"message":        "",
		"goods_ids":      goodsIds,
		"operation_type": consts.OperationInrecycleOperation,
	})
	err := gm.amqpTemplate.Publish(consts.ExchangeGoodsChange,
		consts.ExchangeGoodsChange+"_ROUTING", goodsChangeMsg)
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}
	return nil
}

func (gm *GoodsModel) UnderShopGoods(sellerID int) {
	sqlString := "update es_goods set market_enable = 0 where seller_id = ? "
	if gm.ExecuteSql(sqlString, sellerID) == -1 {
		log.Println("商品更新失败")
	}

	//发送商品下架消息
	sqlString = "select goods_id from es_goods where seller_id = ?"
	rows := gm.QuerySql(sqlString, sellerID)
	defer rows.Close()

	ItemData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
	}
	var goodsIDs []int
	for _, item := range ItemData {
		goodsID := item["goods_id"].(int64)
		goodsIDs = append(goodsIDs, int(goodsID))
	}
	idStr := sql_utils.InSqlStr(goodsIDs)
	if idStr != "" {
		goodsChangeMsg := rabbitmq.BuildMsg(map[string]interface{}{
			"message":        "店铺关闭",
			"goods_ids":      goodsIDs,
			"operation_type": consts.OperationUnderOperation,
		})
		err := gm.amqpTemplate.Publish(consts.ExchangeGoodsChange,
			consts.ExchangeGoodsChange+"_ROUTING", goodsChangeMsg)
		if err != nil {
			log.Printf("[ERROR] %s\n", err)
		}
	}
}

func (gm *GoodsModel) updateGoodsGrade() {
	goodsList, _ := CreateMemberCommentFactory("").queryGoodsGrade()
	if goodsList != nil {
		for _, goods := range goodsList {
			updateSql := "update es_goods set grade=? where goods_id=?"
			goodsID := goods["goods_id"].(int64)
			originGrade := goods["good_rate"].(float64)
			newGrade := strconv.FormatFloat(
				originGrade*100, 'f', 1, 64)

			if gm.ExecuteSql(updateSql, newGrade, goodsID) == -1 {
				log.Println("更新商品等级失败, 商品ID为:", goodsID)
				break
			}
			cacheKey := fmt.Sprintf("%s_%d", consts.GOODS_GRADE, goodsID)
			rds.Put(cacheKey, newGrade, 0)
			// 发送商品消息变化消息
			goodsChangeMsg := rabbitmq.BuildMsg(map[string]interface{}{
				"message":        "",
				"goods_id":       goodsID,
				"operation_type": consts.OperationUpdateOperation,
			})
			err := gm.amqpTemplate.Publish(consts.ExchangeGoodsChange,
				consts.ExchangeGoodsChange+"_ROUTING", goodsChangeMsg)
			if err != nil {
				log.Printf("[ERROR] %s\n", err)
			}
		}
	}

}
