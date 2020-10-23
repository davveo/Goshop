package sql_utils

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

var (
	sqlNullStringType  = reflect.TypeOf(sql.NullString{})
	sqlNullBoolType    = reflect.TypeOf(sql.NullBool{})
	sqlNullFloat64Type = reflect.TypeOf(sql.NullFloat64{})
	sqlNullInt64ype    = reflect.TypeOf(sql.NullInt64{})
)

func ParseJSON(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		log.Println("ParseJSON.Columns错误", err.Error())
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Println("ParseJSON.Scan查询错误", err.Error())
		}
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}

func ParseToStruct(rows *sql.Rows, to interface{}) error {
	/*
		parse sql rows to struct
	*/
	v := reflect.ValueOf(to)
	if v.Elem().Type().Kind() != reflect.Struct {
		return errors.New("expect a struct")
	}

	var scanDest []interface{}
	columnNames, _ := rows.Columns()

	addrByColumnName := map[string]interface{}{}

	for i := 0; i < v.Elem().NumField(); i++ {
		oneValue := v.Elem().Field(i)
		columnName := v.Elem().Type().Field(i).Tag.Get("json")
		if columnName == "" {
			columnName = oneValue.Type().Name()
		}
		addrByColumnName[columnName] = oneValue.Addr().Interface()
	}
	for _, columnName := range columnNames {
		scanDest = append(scanDest, addrByColumnName[columnName])
	}
	return rows.Scan(scanDest...)
}

func Transfer(dest interface{}) error {
	tp := reflect.TypeOf(dest)
	val := reflect.ValueOf(dest)

	if val.Kind() != reflect.Ptr {
		return errors.New("expect struct Ptr")
	}
	nums := val.Elem().NumField()
	for i := 0; i < nums; i++ {
		tpField := tp.Elem().Field(i)
		valField := val.Elem().Field(i)
		if tpField.Type == sqlNullStringType {
			tsField := valField.Addr().Interface().(*sql.NullString)
			// TODO 有点问题
			valField.SetString(tsField.String)
		} else if tpField.Type == sqlNullBoolType {
			tsField := valField.Addr().Interface().(*sql.NullBool)
			valField.SetBool(tsField.Bool)
		} else if tpField.Type == sqlNullFloat64Type {
			tsField := valField.Addr().Interface().(*sql.NullFloat64)
			valField.SetFloat(tsField.Float64)
		} else if tpField.Type == sqlNullInt64ype {
			tsField := valField.Addr().Interface().(*sql.NullInt64)
			valField.SetInt(tsField.Int64)
		} else if tpField.Type.Kind() == reflect.Int ||
			tpField.Type.Kind() == reflect.Int8 || tpField.Type.Kind() == reflect.Int16 ||
			tpField.Type.Kind() == reflect.Int32 || tpField.Type.Kind() == reflect.Int64 ||
			tpField.Type.Kind() == reflect.Uint || tpField.Type.Kind() == reflect.Uint8 ||
			tpField.Type.Kind() == reflect.Uint16 || tpField.Type.Kind() == reflect.Uint32 ||
			tpField.Type.Kind() == reflect.Uint64 || tpField.Type.Kind() == reflect.Float32 ||
			tpField.Type.Kind() == reflect.Float64 || tpField.Type.Kind() == reflect.String ||
			tpField.Type.Kind() == reflect.Slice || tpField.Type.Kind() == reflect.Struct {

		} else {
			return errors.New("unsupport sqlType")
		}
	}
	return nil
}

func GetCountSql(origin string) string {
	end := strings.Index(origin, "from")
	start := strings.Index(origin, "select") + 6
	return strings.ReplaceAll(origin, origin[start:end], " count(*) ")
}

func Like(field, value string, obscure bool) string {
	if obscure {
		return fmt.Sprintf(" and %s like '%s'", field, "%"+value+"%")
	}
	return fmt.Sprintf(" and %s like '%s'", field, value)
}

func LimitOffset(limit, offset int) string {
	if offset == 0 {
		offset = 20 // 默认20
	}
	return fmt.Sprintf(" limit %d, %d", limit-1, offset)
}

func OrderBy(field, order string) string {
	if order == "" {
		order = "desc"
	}
	return fmt.Sprintf(" order by %s %s ", field, order)
}

func Count(sql string, db *sql.DB) (rows int64) {
	err := db.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}
	return rows
}

type Builder struct {
	sqlString bytes.Buffer
	BaseSql   string
}

func (b *Builder) Where(col string, val interface{}, op string) *Builder {
	exp := genExp(col, val, op)
	if b.sqlString.String() == "" {
		b.sqlString.WriteString(fmt.Sprintf("where %s", exp))
	} else {
		b.sqlString.WriteString(fmt.Sprintf(" and %s", exp))
	}
	return b
}

func (b *Builder) OrderBy(field, order string) *Builder {
	if order == "" {
		order = "desc"
	}
	b.sqlString.WriteString(fmt.Sprintf(" order by %s %s ", field, order))
	return b
}

func (b *Builder) LimitOffset(limit, offset int) *Builder {
	if offset == 0 {
		offset = 20 // 默认20
	}
	b.sqlString.WriteString(fmt.Sprintf(" limit %d, %d", limit-1, offset))
	return b
}

func SqlCountString(sqlString string) string {
	end := strings.Index(sqlString, "from")
	start := strings.Index(sqlString, "select") + 6
	return strings.ReplaceAll(sqlString, sqlString[start:end], " count(*) ")
}

func (b *Builder) ToString() string {
	return fmt.Sprintf("%s %s", b.BaseSql, b.sqlString.String())
}

func genExp(col string, val interface{}, op string) string {
	var sExpr string
	switch val.(type) {
	case int, int8, int16, int32, int64:
		sExpr = fmt.Sprintf("%s %s %d", col, op, val)
	case float32, float64:
		sExpr = fmt.Sprintf("%s %s %f", col, op, val)
	case bool:
		var newVal int
		if val == true {
			newVal = 1
		} else {
			newVal = 0
		}
		sExpr = fmt.Sprintf("%s %s %d", col, op, newVal)
	case string:
		newVal, _ := val.(string)
		if op == "like" {
			sExpr = fmt.Sprintf("%s %s '%s'", col, op, "%"+newVal+"%")
		} else {
			sExpr = fmt.Sprintf("%s %s '%s'", col, op, val)
		}

	}

	return deleteExtraSpace(sExpr)
}

func deleteExtraSpace(s string) string {
	s1 := strings.Replace(s, "	", " ", -1)
	regString := "\\s{2,}"
	reg, _ := regexp.Compile(regString)
	s2 := make([]byte, len(s1))
	copy(s2, s1)
	spcIndex := reg.FindStringIndex(string(s2))
	for len(spcIndex) > 0 {
		s2 = append(s2[:spcIndex[0]+1], s2[spcIndex[1]:]...)
		spcIndex = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}

func GetInSql([]int) string {

}
