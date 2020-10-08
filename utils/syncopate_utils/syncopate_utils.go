package syncopate_utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	table = []string{
		"es_sss_order_data",
		"es_sss_order_goods_data",
		"es_sss_refund_data",
		"es_sss_shop_pv",
		"es_sss_goods_pv",
	}
)

type SyncopateUtil struct {
}

func (su *SyncopateUtil) Init(db *sql.DB) {
	var (
		count     int
		sqlString = "select count(0) from es_sss_order_data where create_time > '%s' and  create_time <  '%s'"
	)
	for i := 2015; i < time.Now().Year(); i++ {
		stringIndex := strconv.Itoa(i)
		su.Drop(stringIndex, db)
		yearTime := su.getYearTime(stringIndex)
		err := db.QueryRow(fmt.Sprintf(sqlString, yearTime[0], yearTime[1])).Scan(&count)
		if err != nil {
			log.Println("sql.count 错误", err.Error())
		}
		if count > 0 {
			su.SyncopateTable(stringIndex, db)
		}
	}
}

func (su *SyncopateUtil) HandleSql(year string, sql string) string {
	// 涉及到分表, 所以需要做替换
	// 原表 es_sss_order_data
	// 替换为 es_sss_order_data_2019
	if year == "" || sql == "" {
		return ""
	}
	sql = strings.ToLower(sql)

	return su.replaceTable(year, sql)
}

func (su *SyncopateUtil) replaceTable(year, sql string) string {
	for i := 0; i < len(table); i++ {
		sql = strings.ReplaceAll(sql, table[i], table[i]+"_"+year)
	}
	return sql
}

func (su *SyncopateUtil) CreateTable(year string, db *sql.DB) {
	for _, tb := range table {
		_, _ = db.Exec(fmt.Sprintf("create table %s select *from %s where 1=0", tb+"_"+year, tb))
	}
}

// 切分表
func (su *SyncopateUtil) SyncopateTable(year string, db *sql.DB) {
	yearTime := su.getYearTime(year)
	su.Drop(year, db)

	for _, tb := range table {
		if tb == "es_sss_shop_pv" || tb == "es_sss_goods_pv" {
			_, _ = db.Exec(fmt.Sprintf("create table %s_%s like %s", tb, year, tb))
			_, _ = db.Exec(fmt.Sprintf("insert into %s_%s select * from %s where vs_year = %s", tb, year, tb, year))
		} else {
			_, _ = db.Exec(fmt.Sprintf("create table %s_%s like %s", tb, year, tb))
			_, _ = db.Exec(fmt.Sprintf("insert into %s_%s select * from %s where create_time >= %s and create_time < %s",
				tb, year, tb, yearTime[0], yearTime[1]))
		}
	}

}

func (su *SyncopateUtil) Drop(year string, db *sql.DB) {
	for _, tb := range table {
		db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS '%s_%s'", tb, year))
	}
}

func (su *SyncopateUtil) getYearTime(year string) []string {
	now := int(time.Now().Unix())
	theYearBefore := int(time.Now().AddDate(
		-1, 0, 0).Unix())

	return []string{strconv.Itoa(theYearBefore), strconv.Itoa(now)}
}

func (su *SyncopateUtil) CreateCurrentTable(db *sql.DB) {
	year := strconv.Itoa(time.Now().Year())
	su.SyncopateTable(year, db)
}
