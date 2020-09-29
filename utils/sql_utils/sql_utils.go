package sql_utils

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
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
