package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
	"time"
)

func CreatePinTuanFactory(sqlType string) *PinTuanModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
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
	promotionName, okPromotionName := params["name"].(string)

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
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
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

func (ptm *PinTuanModel) Get(status string) []map[string]interface{} {
	sqlString := "select * from es_pintuan where status = ?"

	if status == consts.UNDERWAY {
		now := time.Now().Unix()
		extra := fmt.Sprintf(" and start_time > %d  and end_time <  %d", now, now)
		sqlString += extra
	}

	rows := ptm.QuerySql(sqlString, status)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}
	return tableData
}
