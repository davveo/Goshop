package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"fmt"
	"log"
)

func CreateExchangeFactory(sqlType string) *ExchangeModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &ExchangeModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("ExchangeModel工厂初始化失败")
	return nil
}

type ExchangeModel struct {
	*BaseModel
}

func (ex *ExchangeModel) deleteByGoods(goodsId int) {
	exchange, _ := ex.getModelByGoods(goodsId)
	if exchange == nil {
		return
	}

	ex.delete(exchange["exchange_id"].(int))
}

func (ex *ExchangeModel) delete(exchangeID int) {
	//删除数据库信息

	sqlString := "delete from es_exchange where exchange_id = ?"
	if ex.ExecuteSql(sqlString, exchangeID) == -1 {
		log.Println("删除数据失败")
	}
	//删除活动商品对照表的关系
	CreatePromotionGoodsFactory("").Delete(exchangeID, consts.PromotionExchange)

	//删除缓存
	rds.Remove(fmt.Sprintf("%s_%d", consts.STORE_ID_EXCHANGE_KEY, exchangeID))
}

func (ex *ExchangeModel) getModelByGoods(goodsId int) (map[string]interface{}, error) {
	sqlString := "select * from es_exchange where goods_id = ? "
	rows := ex.QuerySql(sqlString, goodsId)
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
