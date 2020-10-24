package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"log"
)

func CreateMemberCommentFactory(sqlType string) *MemberCommentModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &MemberCommentModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type MemberCommentModel struct {
	*BaseModel
}

func (mcm *MemberCommentModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_member_comment where status = 1")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	sqlString.WriteString(" order by create_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := mcm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, mcm.count()
}

func (mcm *MemberCommentModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_member_comment;"
	)

	err := mcm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (mcm *MemberCommentModel) queryGoodsGrade() ([]map[string]interface{}, error) {
	sqlString := " select goods_id, sum( CASE grade WHEN '" + consts.CommentGradeGood +
		"' THEN 1 ELSE 0  END ) /count(*) good_rate " +
		" from es_member_comment where status = 1 group by goods_id"

	rows := mcm.QuerySql(sqlString)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, err
	}
	return tableData, nil
}

func (mcm *MemberCommentModel) getModel(commentID int) (map[string]interface{}, error){
	sqlString := "select * from es_member_comment where comment_id = ?"
	rows := mcm.QuerySql(sqlString, commentID)
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

func (mcm *MemberCommentModel) getGoodsCommentCount(goodsID int) (rows int64) {
	sqlString :=  "SELECT COUNT(0) FROM es_member_comment where goods_id = ? and status = 1"
	err := mcm.QueryRow(sqlString).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}
	return rows
}
