package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"errors"
	"fmt"
	"log"
)

func CreateMemberAskFactory(sqlType string) *MemberAskModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &MemberAskModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type MemberAskModel struct {
	*BaseModel
}

func (mam *MemberAskModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_member_ask")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	keyword, okKeyword := params["keyword"].(string)
	goodsName, okGoodsName := params["goodsName"].(string)
	sellerId, okSellerId := params["seller_id"].(string)
	memberName, okMemberName := params["member_name"].(string)
	content, okContent := params["content"].(string)
	authStatus, okAuthStatus := params["auth_status"].(string)
	anonymous, okAnonymous := params["anonymous"].(string)
	replyStatus, okReplyStatus := params["reply_status"].(string)
	memberId, okMemberId := params["member_id"].(string)
	status, okStatus := params["status"].(string)

	endTime, okEndTime := params["end_time"].(string)
	startTime, okStartTime := params["start_time"].(string)

	if status != "" && okStatus {
		sqlString.WriteString(fmt.Sprintf(" where status = %s", status))
	}
	//按关键字查询
	if keyword != "" && okKeyword {
		sqlString.WriteString(fmt.Sprintf(" and (content like '%s' or goods_name like '%s' or member_name like '%s')",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%"))
	}

	//按会员ID查询
	if memberId != "" && okMemberId {
		sqlString.WriteString(fmt.Sprintf(" and member_id = %s", memberId))
	}

	//按商家回复状态查询
	if replyStatus != "" && okReplyStatus {
		sqlString.WriteString(fmt.Sprintf(" and reply_status = %s", replyStatus))
	}

	//按商家ID查询
	if sellerId != "" && okSellerId {
		sqlString.WriteString(fmt.Sprintf(" and seller_id = %s", sellerId))
	}

	//按商品名称查询
	if goodsName != "" && okGoodsName {
		sqlString.WriteString(fmt.Sprintf(" and goods_name like '%s' ", "%"+goodsName+"%"))
	}

	//按会员名称查询
	if memberName != "" && okMemberName {
		sqlString.WriteString(fmt.Sprintf(" and member_name like '%s' ", "%"+memberName+"%"))
	}
	//按咨询内容查询
	if content != "" && okContent {
		sqlString.WriteString(fmt.Sprintf(" and content like '%s' ", "%"+content+"%"))
	}

	//按审核状态查询
	if authStatus != "" && okAuthStatus {
		sqlString.WriteString(fmt.Sprintf(" and auth_status = %s", authStatus))
	}

	//按咨询时间-起始时间查询
	if startTime != "" && okStartTime {
		sqlString.WriteString(fmt.Sprintf(" and create_time >= %s", startTime))
	}

	//按咨询时间-结束时间查询
	if endTime != "" && okEndTime {
		sqlString.WriteString(fmt.Sprintf(" and create_time <= %s", endTime))
	}

	//按匿名状态查询
	if anonymous != "" && okAnonymous {
		sqlString.WriteString(fmt.Sprintf(" and anonymous = %s", anonymous))
	}

	sqlString.WriteString(" order by create_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := mam.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, mam.count(status)
}

func (mam *MemberAskModel) count(status string) (rows int64) {
	var (
		sql = "select count(*) from es_member_ask"
	)

	if status != "" {
		sql += fmt.Sprintf(" where status = %s", status)
	}

	err := mam.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (mam *MemberAskModel) delete(askId int) (success bool) {
	sql := "update es_member_ask set status = ? where ask_id = ?"
	success = mam.ExecuteSql(sql, consts.DELETED, askId) < 0

	// TODO 同时删除咨询问题的回复和发送的站内消息
	// this.askReplyManager.deleteByAskId(askId);
	// this.askMessageManager.deleteByAskId(askId);
	return
}

func (mam *MemberAskModel) Reply(replyContent string, askId int) error {
	if replyContent == "" {
		return errors.New("回复内容不能为空")
	}

	if len(replyContent) < 3 || len(replyContent) > 120 {
		return errors.New("回复内容应在3到120个字符之间")
	}
	return nil
}
