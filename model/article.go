package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"errors"
	"log"
)

func CreateArticleFactory(sqlType string) *ArticleModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &ArticleModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("brandModel工厂初始化失败")
	return nil
}

type ArticleModel struct {
	*BaseModel
}

func (am *ArticleModel) GetModel(articleID int) (map[string]interface{}, error) {
	rows := am.QuerySql("select * from es_article where article_id = ?", articleID)
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

func (am *ArticleModel) List(params map[string]string) (map[string]interface{}, int64) {
	return nil, 0
}

func (am *ArticleModel) Add(params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func (am *ArticleModel) Edit(params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func (am *ArticleModel) Delete(articleId int) error {
	if am.ExecuteSql("delete from es_article where article_id = ?", articleId) == -1 {
		return errors.New("删除数据失败")
	}
	return nil
}
