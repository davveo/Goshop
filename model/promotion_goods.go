package model

import (
	"Goshop/utils/yml_config"
	"log"
)

func CreatePromotionGoodsFactory(sqlType string) *PromotionGoodsModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &PromotionGoodsModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("ExchangeModel工厂初始化失败")
	return nil
}

type PromotionGoodsModel struct {
	*BaseModel
}

func (pgm *PromotionGoodsModel) Delete(activityId int, promotionType string) {

	sqlString := "DELETE FROM es_promotion_goods WHERE activity_id=? and promotion_type= ? "
	if pgm.ExecuteSql(sqlString, activityId, promotionType) == -1 {
		log.Println("删除数据失败")
	}

}
