package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateCouponFactory(sqlType string) *CouponModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &CouponModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type CouponModel struct {
	*BaseModel
}

func (cm *CouponModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_coupon")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	keyword, okKeyword := params["keyword"].(string)
	endTime, okEndTime := params["end_time"].(string)
	sellerId, okSellerId := params["seller_id"].(string)
	startTime, okStartTime := params["start_time"].(string)

	if sellerId != "" && sellerId != "0" && okSellerId {
		sqlString.WriteString(fmt.Sprintf(" where seller_id = %s", sellerId))
	}

	if startTime != "" && okStartTime {
		sqlString.WriteString(fmt.Sprintf(" and start_time >= %s", startTime))
	}

	if endTime != "" && okEndTime {
		sqlString.WriteString(fmt.Sprintf(" and end_time <= %s", endTime))
	}

	if keyword != "" && okKeyword {
		sqlString.WriteString(fmt.Sprintf(" and title like  '%s'", "%"+keyword+"%"))
	}

	sqlString.WriteString(" order by coupon_id desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := cm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, cm.count()
}

func (cm *CouponModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_pintuan "
	)

	err := cm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
