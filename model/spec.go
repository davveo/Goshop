package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"errors"
	"log"
)

func CreateSpecFactory(sqlType string) *SpecModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &SpecModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type SpecModel struct {
	*BaseModel
}

func (s *SpecModel) GetModel(specID int) (map[string]interface{}, error) {
	sqlString := "select * from es_brand where spec_id = ?"
	rows := s.QuerySql(sqlString, specID)
	defer rows.Close()

	tableData, _ := sql_utils.ParseJSON(rows)
	var tmp map[string]interface{}
	if len(tableData) > 0 {
		tmp = tableData[0]
	}
	return tmp, nil
}

func (s *SpecModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString      bytes.Buffer
		countSqlString string
	)

	sqlString.WriteString("select * from es_specification  where disabled = 1 and seller_id = 0 ")

	pageNo, okPageNo := params["page_no"].(int)
	keyword, okKeyword := params["keyword"].(string)
	pageSize, okPageSize := params["page_size"].(int)

	if keyword != "" && okKeyword {
		sqlString.WriteString(sql_utils.Like("spec_name", keyword, true))
	}

	sqlString.WriteString(sql_utils.OrderBy("spec_id", "desc"))
	countSqlString = sql_utils.GetCountSql(sqlString.String())

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := s.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, sql_utils.Count(countSqlString, s.dbDriverRead)
}

func (s *SpecModel) Add(params map[string]interface{}) (map[string]interface{}, error) {
	var (
		specId    int64
		sqlString string
	)

	sellerId := params["seller_id"].(string)
	disabled := params["disabled"].(int)
	specName := params["spec_name"].(string)
	specMemo := params["spec_memo"].(string)

	//如果是管理端添加的规格，则验证管理端的对个名称是否重复
	if sellerId == "0" {
		sqlString = "select * from es_specification  where disabled = 1 and seller_id = 0 and spec_name = ? "
		rows := s.QuerySql(sqlString, specName)
		defer rows.Close()

		tableData, _ := sql_utils.ParseJSON(rows)
		if len(tableData) > 0 {
			return nil, errors.New("规格名称重复")
		}
	}

	sqlString = "insert into `es_specification` (`spec_name`, `disabled`, `spec_memo`, `seller_id`) values (?,?,?,?)"

	if specId = s.LastInsertId(sqlString, specName, disabled, specMemo, sellerId); specId == -1 {
		return nil, errors.New("插入规格失败")
	}

	params["spec_id"] = specId
	params["disabled"] = 1
	return params, nil
}

func (s *SpecModel) Edit(params map[string]interface{}) (map[string]interface{}, error) {
	var (
		sqlString string
	)

	specId := params["spec_id"].(int)
	disabled := params["disabled"].(int)
	sellerId := params["seller_id"].(string)
	specName := params["spec_name"].(string)
	specMemo := params["spec_memo"].(string)

	spec, _ := s.GetModel(specId)
	if spec == nil {
		return nil, errors.New("规格不存在")
	}

	sqlString = "select * from es_specification  " +
		"where disabled = 1 and seller_id = 0 and spec_name = ? and spec_id!=? "

	rows := s.QuerySql(sqlString, specName, specId)
	defer rows.Close()

	specList, _ := sql_utils.ParseJSON(rows)

	if len(specList) > 0 {
		return nil, errors.New("规格名称重复")
	}
	sqlString = "update  es_specification set `spec_name` = ?, `disabled` = ?, `spec_memo` = ?, `seller_id` = ? where spec_id = ?"
	if affected := s.ExecuteSql(sqlString, specName, disabled, specMemo, sellerId); affected == -1 {
		return nil, errors.New("更新失败")
	}

	return params, nil
}

func (s *SpecModel) Delete(specIds []int) error {
	var (
		err       error
		sqlString string
	)
	idsStr := sql_utils.InSqlStr(specIds)

	//查看是否已经有分类绑定了该规格
	sqlString = "select * from es_category_spec where spec_id in (" + idsStr + ")"

	rows := s.QuerySql(sqlString)
	defer rows.Close()

	specList, err := sql_utils.ParseJSON(rows)

	if len(specList) > 0 || err != nil {
		return errors.New("有分类已经绑定要删除的规格，请先解绑分类规格")
	}
	sqlString = " update es_specification set disabled = 0 where spec_id in (" + idsStr + ")"
	if affected := s.ExecuteSql(sqlString); affected == -1 {
		return errors.New("删除规格失败")
	}
	return nil
}
