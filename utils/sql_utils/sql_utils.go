package sql_utils

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
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
