package model

import (
	"Goshop/global/variable"
	"Goshop/utils/yml_config"
	"encoding/json"
	"fmt"
	"log"
)

type Setting struct {
	SiteName      string `json:"site_name"`
	Title         string `json:"title"`
	Keywords      string `json:"keywords"`
	Descript      string `json:"descript"`
	Siteon        int    `json:"siteon"`
	CloseReson    string `json:"close_reson"`
	Logo          string `json:"logo"`
	GlobalAuthKey string `json:"global_auth_key"`
	DefaultImg    string `json:"default_img"`
	TestMode      int    `json:"test_mode"`
}

func CreateSettingFactory(sqlType string) *SettingModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &SettingModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("settingModel工厂初始化失败")
	return nil
}

type SettingModel struct {
	*BaseModel
	CfgValue string
}

func (u *SettingModel) Get(group string) *Setting {
	var (
		setting *Setting
		sql     = "select cfg_value from es_settings where cfg_group = ?"
	)

	cacheKey := u.cacheName(variable.SettingsPrefix, group)
	// 从缓存中获取配置
	exist, value := rds.Gain(cacheKey)
	if !exist {
		// 如果没有就从数据库中获取
		rows := u.QuerySql(sql, group)
		if rows != nil {
			for rows.Next() {
				err := rows.Scan(&u.CfgValue)
				if err == nil {
					value = u.CfgValue
					_ = rds.Put(cacheKey, u.CfgValue, 0)
				}
			}
			_ = rows.Close()
		}
	}

	if value != "" {
		_ = json.Unmarshal([]byte(value), &setting)
	}
	return setting
}

func (u *SettingModel) cacheName(prefix string, params ...interface{}) string {
	return fmt.Sprintf(prefix, params...)
}
