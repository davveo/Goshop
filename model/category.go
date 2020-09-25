package model

import (
	"log"
	"orange/utils/yml_config"
)

func CreateCategoryFactory(sqlType string) *CategoryModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &CategoryModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("settingModel工厂初始化失败")
	return nil
}

type BaseCategory struct {
	AdvImage      string `json:"adv_image"`
	AdvImageLink  string `json:"adv_image_link"`
	BrandList     string `json:"brand_list"`
	CategoryID    int    `json:"category_id"`
	CategoryOrder string `json:"category_order"`
	CategoryPath  string `json:"category_path"`
	GoodsCount    string `json:"goods_count"`
	Image         string `json:"image"`
	IsShow        string `json:"is_show"`
	Name          string `json:"name"`
	ParentID      string `json:"parent_id"`
}

type CategoryTree struct {
	BaseCategory
	Children []CategoryTree `json:"children" jpath:"children"`
}

type CategoryModel struct {
	*BaseModel
	BaseCategory
}

func (cm *CategoryModel) List(parentID int) (error, []CategoryTree) {
	return cm.getCategoryTree(parentID)
}

func (cm *CategoryModel) QueryTargetParentTree(parentID int) (err error, allCategory []CategoryTree) {
	var (
		sqlString = "select c.* from es_category c  where c.parent_id = ? order by c.category_order asc"
	)

	rows := cm.QuerySql(sqlString, parentID)
	defer rows.Close()

	if rows != nil {
		for rows.Next() {
			// 与数据库列对其 坑!!!
			err := rows.Scan(
				&cm.CategoryID, &cm.Name, &cm.ParentID, &cm.CategoryPath,
				&cm.GoodsCount, &cm.CategoryOrder, &cm.Image, &cm.IsShow,
				&cm.AdvImage, &cm.AdvImageLink)
			if err == nil {
				// TODO image字段存在问题
				allCategory = append(
					allCategory, CategoryTree{
						BaseCategory: cm.BaseCategory, Children: nil,
					})
			} else {
				log.Println("sql查询错误", err.Error())
			}
		}
		_ = rows.Close()
	}

	return err, allCategory
}

func (cm *CategoryModel) getCategoryTree(parentID int) (error, []CategoryTree) {
	var categoryTree []CategoryTree
	_, categoryList := cm.QueryTargetParentTree(parentID)

	for _, v := range categoryList {
		_, trees := cm.getCategoryTree(v.CategoryID)
		v.Children = trees
		categoryTree = append(categoryTree, v)
	}
	return nil, categoryTree
}
