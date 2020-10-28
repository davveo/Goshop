package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"errors"
	"log"
)

func CreateSpecValuesFactory(sqlType string) *SpecValuesModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &SpecValuesModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("SpecValuesModel工厂初始化失败")
	return nil
}

type SpecValuesModel struct {
	*BaseModel
}

func (svm *SpecValuesModel) GetModel(specValueId int) (map[string]interface{}, error) {
	sqlString := "select * from es_spec_values where spec_value_id = ?"
	rows := svm.QuerySql(sqlString, specValueId)
	defer rows.Close()

	tableData, _ := sql_utils.ParseJSON(rows)
	var tmp map[string]interface{}
	if len(tableData) > 0 {
		tmp = tableData[0]
	}
	return tmp, nil
}

func (svm *SpecValuesModel) ListBySpecId(specId, permission int) []map[string]interface{} {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_spec_values where spec_id = ? ")

	if consts.PermissionADMIN == permission {
		sqlString.WriteString("  and seller_id = 0")
	}
	rows := svm.QuerySql(sqlString.String(), specId)
	defer rows.Close()
	tableData, _ := sql_utils.ParseJSON(rows)
	return tableData
}

//func (svm *SpecValuesModel) Add(params map[string]interface{}) (map[string]interface{}, error) {
//	var (
//		specId    int64
//		sqlString string
//	)
//
//	sellerId := params["seller_id"].(string)
//	disabled := params["disabled"].(int)
//	specName := params["spec_name"].(string)
//	specMemo := params["spec_memo"].(string)
//
//	sqlString = "insert into `es_spec_values` (`spec_id`, `spec_value`, `seller_id`, `spec_name`) values (?,?,?,?)"
//
//	if specId = svm.LastInsertId(sqlString, specName, disabled, specMemo, sellerId); specId == -1 {
//		return nil, errors.New("插入规格失败")
//	}
//
//	params["spec_id"] = specId
//	return params, nil
//}

//
//func (svm *SpecValuesModel) Edit(params map[string]interface{}) (map[string]interface{}, error) {
//	var (
//		sqlString string
//	)
//
//	specId := params["spec_id"].(int)
//	disabled := params["disabled"].(int)
//	sellerId := params["seller_id"].(string)
//	specName := params["spec_name"].(string)
//	specMemo := params["spec_memo"].(string)
//
//	spec, _ := s.GetModel(specId)
//	if spec == nil {
//		return nil, errors.New("规格不存在")
//	}
//
//	sqlString = "select * from es_specification  " +
//		"where disabled = 1 and seller_id = 0 and spec_name = ? and spec_id!=? "
//
//	rows := s.QuerySql(sqlString, specName, specId)
//	defer rows.Close()
//
//	specList, _ := sql_utils.ParseJSON(rows)
//
//	if len(specList) > 0 {
//		return nil, errors.New("规格名称重复")
//	}
//	sqlString = "update  es_specification set `spec_name` = ?, `disabled` = ?, `spec_memo` = ?, `seller_id` = ? where spec_id = ?"
//	if affected := s.ExecuteSql(sqlString, specName, disabled, specMemo, sellerId); affected == -1 {
//		return nil, errors.New("更新失败")
//	}
//
//	return params, nil
//}

func (svm *SpecValuesModel) SaveSpecValue(specId int, valueList []string) ([]map[string]interface{}, error) {
	var ret []map[string]interface{}
	spec, err := svm.GetModel(specId)
	if spec == nil || err != nil {
		return nil, errors.New("所属规格不存在")
	}
	sqlString := "delete from es_spec_values where spec_id=? and seller_id=0"
	if affected := svm.ExecuteSql(sqlString, specId); affected == -1 {
		return nil, errors.New("更新失败")
	}
	for _, value := range valueList {
		var insertId int64
		if len(value) > 50 {
			return nil, errors.New("规格值为1到50个字符之间")
		}
		sqlString = "insert into `es_spec_values` (`spec_id`, `spec_value`, `seller_id`, `spec_name`) values (?,?,?,?)"
		if insertId = svm.LastInsertId(sqlString, specId, value, 0, spec["spec_name"]); insertId == -1 {
			return nil, errors.New("更新失败")
		}
		ret = append(ret, map[string]interface{}{
			"spec_value_id": insertId,
			"spec_id":       specId,
			"spec_value":    value,
			"spec_name":     spec["spec_name"],
			"seller_id":     0,
		})
	}
	return ret, nil
}
