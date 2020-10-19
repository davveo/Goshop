package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateMemberAskReplyFactory(sqlType string) *AskReplyModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AskReplyModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type AskReplyModel struct {
	*BaseModel
}

func (arm *AskReplyModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_ask_reply where is_del = 'NORMAL'")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	askId, okAskId := params["ask_id"].(string)
	keyword, okKeyword := params["keyword"].(string)
	memberName, okMemberName := params["member_name"].(string)
	content, okContent := params["content"].(string)
	authStatus, okAuthStatus := params["auth_status"].(string)
	anonymous, okAnonymous := params["anonymous"].(string)
	replyStatus, okReplyStatus := params["reply_status"].(string)
	memberId, okMemberId := params["member_id"].(string)
	id, okId := params["id"].(string)

	endTime, okEndTime := params["end_time"].(string)
	startTime, okStartTime := params["start_time"].(string)

	//按会员商品咨询ID查询
	if askId != "" && okAskId {
		sqlString.WriteString(fmt.Sprintf(" where ask_id = %s", askId))
	}
	//按关键字查询
	if keyword != "" && okKeyword {
		sqlString.WriteString(fmt.Sprintf(" and (content like '%s' or member_name like '%s')", "%"+keyword+"%", "%"+keyword+"%"))
	}

	//按会员名称查询
	if memberName != "" && okMemberName {
		sqlString.WriteString(fmt.Sprintf(" and member_name like '%s' ", "%"+memberName+"%"))
	}

	//按回复内容查询
	if content != "" && okContent {
		sqlString.WriteString(fmt.Sprintf(" and content like '%s' ", "%"+content+"%"))
	}

	//按审核状态查询
	if authStatus != "" && okAuthStatus {
		sqlString.WriteString(fmt.Sprintf(" and auth_status = %s", authStatus))
	}

	//按回复时间-起始时间查询
	if startTime != "" && okStartTime {
		sqlString.WriteString(fmt.Sprintf(" and reply_time >= %s", startTime))
	}

	//按回复时间-结束时间查询
	if endTime != "" && okEndTime {
		sqlString.WriteString(fmt.Sprintf(" and reply_time <= %s", endTime))
	}

	//按匿名状态查询
	if anonymous != "" && okAnonymous {
		sqlString.WriteString(fmt.Sprintf(" and anonymous = %s", anonymous))
	}

	//按回复状态查询
	if replyStatus != "" && okReplyStatus {
		sqlString.WriteString(fmt.Sprintf(" and reply_status = %s", replyStatus))
	}

	//按会员id查询
	if memberId != "" && okMemberId {
		sqlString.WriteString(fmt.Sprintf(" and member_id = %s", memberId))
	}

	//排除某条回复（一般用于商品详情页面获取咨询回复）
	if id != "" && okId {
		sqlString.WriteString(fmt.Sprintf(" and id = %s", id))
	}

	sqlString.WriteString(" order by reply_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := arm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, arm.count()
}

func (arm *AskReplyModel) count() (rows int64) {
	var (
		sql = "select * from es_ask_reply where is_del = 'NORMAL'"
	)

	err := arm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
