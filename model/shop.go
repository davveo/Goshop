package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreatGoshopFactory(sqlType string) *ShopModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &ShopModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("shopModel工厂初始化失败")
	return nil
}

type ShopModel struct {
	*BaseModel
	ID                      sql.NullString `json:"-"`
	MemberId                sql.NullString `json:"member_id"`
	MemberName              sql.NullString `json:"member_name"`
	ShopName                sql.NullString `json:"shop_name"`
	ShopDisable             sql.NullString `json:"shop_disable"`
	ShopCreatetime          sql.NullString `json:"shop_createtime"`
	ShopEdtime              sql.NullString `json:"shop_endtime"`
	ShopId                  sql.NullString `json:"shop_id"`
	ShopProvinceId          sql.NullString `json:"shop_province_id"`
	ShopCityId              sql.NullString `json:"shop_city_id"`
	ShopCountyId            sql.NullString `json:"shop_county_id"`
	ShopTownId              sql.NullString `json:"shop_town_id"`
	ShopProvince            sql.NullString `json:"shop_province"`
	ShopCity                sql.NullString `json:"shop_city"`
	ShopCounty              sql.NullString `json:"shop_county"`
	ShopTown                sql.NullString `json:"shop_town"`
	ShopAdd                 sql.NullString `json:"shop_add"`
	CompanyName             sql.NullString `json:"company_name"`
	CompanyAddress          sql.NullString `json:"company_address"`
	CompanyPhone            sql.NullString `json:"company_phone"`
	CompanyEmail            sql.NullString `json:"company_email"`
	EmployeeNum             sql.NullInt64  `json:"employee_num"`
	RegMoney                sql.NullString `json:"reg_money"`
	LinkName                sql.NullString `json:"link_name"`
	LinkPhone               sql.NullString `json:"link_phone"`
	LegalName               sql.NullString `json:"legal_name"`
	LegalId                 sql.NullString `json:"legal_id"`
	LegalImg                sql.NullString `json:"legal_img"`
	LicenseNum              sql.NullString `json:"license_num"`
	LicenseProvinceId       sql.NullString `json:"license_province_id"`
	LicenseCityId           sql.NullString `json:"license_city_id"`
	LicenseCountyId         sql.NullString `json:"license_county_id"`
	LicenseTownId           sql.NullString `json:"license_town_id"`
	LicenseProvince         sql.NullString `json:"license_province"`
	LicenseCity             sql.NullString `json:"license_city"`
	LicenseCounty           sql.NullString `json:"license_county"`
	LicenseTown             sql.NullString `json:"license_town"`
	LicenseAdd              sql.NullString `json:"license_add"`
	EstablishDate           sql.NullString `json:"establish_date"`
	LicenceStart            sql.NullString `json:"licence_start"`
	LicenceEnd              sql.NullString `json:"licence_end"`
	Scope                   sql.NullString `json:"scope"`
	LicenceImg              sql.NullString `json:"licence_img"`
	OrganizationCode        sql.NullString `json:"organization_code"`
	CodeImg                 sql.NullString `json:"code_img"`
	TaxesImg                sql.NullString `json:"taxes_img"`
	BankAccountName         sql.NullString `json:"bank_account_name"`
	BankNumber              sql.NullString `json:"bank_number"`
	BankName                sql.NullString `json:"bank_name"`
	BankProvinceId          sql.NullString `json:"bank_province_id"`
	BankCityId              sql.NullString `json:"bank_city_id"`
	BankCountyId            sql.NullString `json:"bank_county_id"`
	BankTownId              sql.NullString `json:"bank_town_id"`
	BankProvince            sql.NullString `json:"bank_province"`
	BankCity                sql.NullString `json:"bank_city"`
	BankCounty              sql.NullString `json:"bank_county"`
	BankTown                sql.NullString `json:"bank_town"`
	BankImg                 sql.NullString `json:"bank_img"`
	TaxesCertificateNum     sql.NullString `json:"taxes_certificate_num"`
	TaxesDistinguishNum     sql.NullString `json:"taxes_distinguish_num"`
	TaxesCertificateImg     sql.NullString `json:"taxes_certificate_img"`
	GoodsManagementCategory sql.NullString `json:"goods_management_category"`
	ShopLevel               sql.NullString `json:"shop_level"`
	ShopLevelApply          sql.NullString `json:"shop_level_apply"`
	StoreSpaceCapacity      sql.NullString `json:"store_space_capacity"`
	ShopLogo                sql.NullString `json:"shop_logo"`
	ShopBanner              sql.NullString `json:"shop_banner"`
	ShopDesc                sql.NullString `json:"shop_desc"`
	ShopRecommend           sql.NullString `json:"shop_recommend"`
	ShopThemeid             sql.NullString `json:"shop_themeid"`
	ShopThemePath           sql.NullString `json:"shop_theme_path"`
	WapThemeid              sql.NullString `json:"wap_themeid"`
	WapThemePath            sql.NullString `json:"wap_theme_path"`
	ShopCredit              sql.NullString `json:"shop_credit"`
	ShopPraiseRate          sql.NullString `json:"shop_praise_rate"`
	ShopDescriptionCredit   sql.NullString `json:"shop_description_credit"`
	ShopServiceCredit       sql.NullString `json:"shop_service_credit"`
	ShopDeliveryCredit      sql.NullString `json:"shop_delivery_credit"`
	ShopCollect             sql.NullString `json:"shop_collect"`
	GoodsNum                sql.NullString `json:"goods_num"`
	ShopQq                  sql.NullString `json:"shop_qq"`
	ShopCommission          sql.NullString `json:"shop_commission"`
	GoodsWarningCount       sql.NullString `json:"goods_warning_count"`
	SelfOperated            sql.NullString `json:"self_operated"`
	Step                    int            `json:"step"`
	OrdinReceiptStatus      int            `json:"ordin_receipt_status"`
	ElecReceiptStatus       int            `json:"elec_receipt_status"`
	TaxReceiptStatus        int            `json:"tax_receipt_status"`
}

func (sm *ShopModel) All() []map[string]interface{} {
	sqlString := "select  s.member_id,s.member_name,s.shop_name,s.shop_disable,s.shop_createtime,s.shop_endtime,sd.* " +
		"from es_shop s  left join es_shop_detail sd on s.shop_id = sd.shop_id   " +
		"where  shop_disable = 'OPEN' order by s.shop_createtime desc"

	rows := sm.QuerySql(sqlString)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}

	return tableData
}

func (sm *ShopModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	shopType := params["shop_type"].(string)
	shopDisable := params["shop_disable"].(string)
	keyword := params["keyword"].(string)
	shop_name := params["shop_name"].(string)

	if shopDisable == "" {
		shopDisable = "OPEN"
	}

	if shopDisable == "ALL" {
		sqlString.WriteString("select  s.member_id,s.member_name,s.shop_name,s.shop_disable,s.shop_createtime," +
			"s.shop_endtime,sd.* from es_shop s  left join es_shop_detail sd on s.shop_id = sd.shop_id  where  shop_disable != 'APPLYING' ")
	} else {
		sqlString.WriteString("select  s.member_id,s.member_name,s.shop_name,s.shop_disable,s.shop_createtime," +
			"s.shop_endtime,sd.* from es_shop s  left join es_shop_detail sd on s.shop_id = sd.shop_id   where  shop_disable = '" + shopDisable + "'")
	}

	if shopType != "" {
		// TODO
	}

	if keyword != "" {
		sqlString.WriteString(fmt.Sprintf("  and (s.shop_name like %s or s.member_name like %s) ",
			"'"+keyword+"'", "'"+keyword+"'"))
	}

	if shop_name != "" {
		sqlString.WriteString(fmt.Sprintf("  and s.shop_name like %s ", "'"+shop_name+"'"))
	}

	sqlString.WriteString(" order by create_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := sm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, sm.count()
}

func (sm *ShopModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_shop"
	)

	err := sm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
