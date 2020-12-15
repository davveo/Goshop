package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateBillMemberFactory(sqlType string) *BillMemberModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &BillMemberModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("BillMemberModel工厂初始化失败")
	return nil
}

type BillMemberModel struct {
	*BaseModel
}

func (bmm *BillMemberModel) List(query map[string]interface{}) ([]map[string]interface{}, int64) {
	var sqlString bytes.Buffer

	pageNo := query["page_no"].(int)
	pageSize := query["page_size"].(int)
	totalId := query["total_id"].(string) // 总结算单id

	sqlString.WriteString("select * from es_bill_member b where 1=1 ")
	if uname, ok := query["uname"].(string); uname != "" && ok {
		sqlString.WriteString(fmt.Sprintf(" and total_id = %s and b.member_name like('%s')",
			totalId, "%"+uname+"%"))
	} else {
		sqlString.WriteString(fmt.Sprintf(" and total_id = %s", totalId))
	}
	sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))

	resSql := sqlString.String()
	rows := bmm.QuerySql(resSql)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}
	return tableData, bmm.count(resSql)
}

func (bmm *BillMemberModel) count(SqlString string) (rows int64) {
	err := bmm.QueryRow(SqlString).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (bmm *BillMemberModel) GetBillMember(billId string) (map[string]interface{}, error) {
	rows := bmm.QuerySql("select * from es_bill_member where id = ?", billId)
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

func (bmm *BillMemberModel) AllDown(billId, memberId string) []map[string]interface{} {
	rows := bmm.QuerySql("select * from es_bill_member where member_id in "+
		"(select member_id from es_distribution where member_id_lv1 =? or member_id_lv2 = ?) "+
		"and total_id = (select total_id from es_bill_member where id = ?)", memberId, memberId, billId)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}
	// 获取下级 分销商集合
	// TODO
	return tableData
}
