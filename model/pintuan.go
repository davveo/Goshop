package model

import (
	"bytes"
	"fmt"
	"log"
	"Eshop/utils/sql_utils"
	"Eshop/utils/yml_config"
	"strconv"
)

func CreatePinTuanFactory(sqlType string) *PinTuanModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &PinTuanModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type PinTuanModel struct {
	*BaseModel
}

func (ptm *PinTuanModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_pintuan ")

	pageNo, okPageNo := params["page_no"].(int)
	status, okStatus := params["status"].(string)
	pageSize, okPageSize := params["page_size"].(int)
	endTime, okEndTime := params["end_time"].(string)
	sellerId, okSellerId := params["seller_id"].(string)
	startTime, okStartTime := params["start_time"].(string)
	promotionName, okPromotionName := params["promotion_name"].(string)

	if sellerId != "" && sellerId != "0" && okSellerId {
		sqlString.WriteString(fmt.Sprintf(" where seller_id = %s", sellerId))
	}

	if startTime != "" && okStartTime {
		sqlString.WriteString(fmt.Sprintf(" and start_time >= %s", startTime))
	}

	if endTime != "" && okEndTime {
		sqlString.WriteString(fmt.Sprintf(" and end_time <= %s", endTime))
	}

	if promotionName != "" && okPromotionName {
		sqlString.WriteString(fmt.Sprintf(" and promotion_name like  '%s'", "%"+promotionName+"%"))
	}

	if status != "" && okStatus {
		sqlString.WriteString(fmt.Sprintf(" and status = %s", status))
	}

	sqlString.WriteString(" order by create_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := ptm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, ptm.count()
}

func (ptm *PinTuanModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_pintuan"
	)

	err := ptm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
