package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/time_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateSeckillFactory(sqlType string) *SeckillModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &SeckillModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type SeckillModel struct {
	*BaseModel
}

func (sm *SeckillModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_seckill")

	pageNo, okPageNo := params["page_no"].(int)
	status, okStatus := params["status"].(string)
	pageSize, okPageSize := params["page_size"].(int)
	endTime, okEndTime := params["end_time"].(string)
	startTime, okstartTime := params["start_time"].(string)
	seckillName, okSeckillName := params["seckill_name"].(string)
	deleteStatus, okDeleteStatus := params["delete_status"].(string)

	startOfDay := time_utils.GetToDayOfStart()
	endOfDay := time_utils.GetToDayOfEnd()

	sqlString.WriteString(" order by start_day desc")

	if deleteStatus != "" && okDeleteStatus {
		sqlString.WriteString(fmt.Sprintf(" where delete_status = '%s'", deleteStatus))
	}

	if status != "" && okStatus && status != "ALL" {
		if consts.SeckillEditing == status {
			sqlString.WriteString(fmt.Sprintf(" and seckill_status = '%s'", status))
		} else if consts.SeckillRelease == status {
			sqlString.WriteString(fmt.Sprintf(" and seckill_status = '%s' and start_day > '%s'", status, endOfDay))
		} else if consts.SeckillOpen == status {
			sqlString.WriteString(fmt.Sprintf(" and seckill_status = '%s' and start_day >= '%s' and start_day <= '%s'",
				consts.SeckillRelease, startOfDay, endOfDay))
		} else if consts.SeckillClosed == status {
			sqlString.WriteString(fmt.Sprintf(" and seckill_status = '%s' and start_day < '%s'",
				consts.SeckillRelease, startOfDay))
		}
	}

	if startTime != "" && okstartTime {
		sqlString.WriteString(fmt.Sprintf(" and start_day >= '%s'", startTime))
	}

	if endTime != "" && okEndTime {
		sqlString.WriteString(fmt.Sprintf(" and start_day <= %s", endTime))
	}

	if seckillName != "" && okSeckillName {
		sqlString.WriteString(fmt.Sprintf(" and seckill_name like  '%s'", "%"+seckillName+"%"))
	}

	sqlString.WriteString(" order by start_day desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := sm.QuerySql(sqlString.String())
	defer rows.Close()

	seckillItemList, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	for _, seckillItem := range seckillItemList {
		var rangeList []int
		seckillId := seckillItem["seckill_id"].(int)
		seckillStatus := seckillItem["seckill_status"].(string)
		seckillRangeList := CreateSeckillRangeFactory("").getList(seckillId)
		if seckillRangeList == nil {
			continue
		}
		for _, seckillRangeitem := range seckillRangeList {
			rangeTime := seckillRangeitem["range_time"].(int)
			rangeList = append(rangeList, rangeTime)
		}
		if seckillStatus != "" {
			//如果状态是已发布状态，则判断该活动是否已开始或者已结束
			seckillItem["seckill_status_text"] = consts.SeckillMap[seckillStatus]
			if seckillStatus == consts.SeckillRelease {
				startDay := seckillItem["start_day"].(int64)
				if time_utils.StartOfDay() <= startDay && time_utils.EndOfDay() > startDay {
					seckillItem["seckill_status_text"] = "已开启"
				} else if startDay < time_utils.EndOfDay() {
					seckillItem["seckill_status_text"] = "已关闭"
				}
			}
		}
		seckillItem["range_list"] = rangeList
	}

	return seckillItemList, sm.count()
}

func (sm *SeckillModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_seckill"
	)

	err := sm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
