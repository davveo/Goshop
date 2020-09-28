package model

import (
	"log"
	"Eshop/utils/sql_utils"
	"Eshop/utils/yml_config"
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
	AdvImage      string `sql:"adv_image" json:"adv_image"`
	AdvImageLink  string `sql:"adv_image_link" json:"adv_image_link"`
	BrandList     string `sql:"brand_list" json:"brand_list"`
	CategoryID    int    `sql:"category_id" json:"category_id"`
	CategoryOrder string `sql:"category_order" json:"category_order"`
	CategoryPath  string `sql:"category_path" json:"category_path"`
	GoodsCount    string `sql:"goods_count" json:"goods_count"`
	Image         string `sql:"image" json:"image"`
	IsShow        string `sql:"is_show" json:"is_show"`
	Name          string `sql:"name" json:"name"`
	ParentID      string `sql:"parent_id" json:"parent_id"`
}

type CategoryTree struct {
	BaseCategory
	Children []CategoryTree `json:"children"`
}

type CategoryModel struct {
	*BaseModel
	BaseCategory
}

func (cm *CategoryModel) List(parentID int) (error, []CategoryTree) {
	return cm.getCategoryTree(parentID)
}

func (cm *CategoryModel) QueryTargetParentTree(parentID int) (allCategory []CategoryTree) {
	var (
		sqlString = "select c.* from es_category c  where c.parent_id = ? order by c.category_order asc"
	)

	rows := cm.QuerySql(sqlString, parentID)
	defer rows.Close()

	if rows != nil {
		for rows.Next() {
			// TODO 与数据库列对其 坑!!!
			// rows.Scan 参数的顺序很重要, 需要和查询的结果的column对应.
			// 例如 “SELECT * From user where age >=20 AND age < 30”
			// 查询的行的 column 顺序是 “id, name, age” 和插入操作顺序相同,
			// 因此 rows.Scan 也需要按照此顺序 rows.Scan(&id, &name, &age), 不然会造成数据读取的错位.
			// https://www.cnblogs.com/hanyouchun/p/6708037.html
			baseCategory := BaseCategory{}
			_ = sql_utils.ParseToStruct(rows, &baseCategory)
			allCategory = append(allCategory, CategoryTree{BaseCategory: baseCategory, Children: nil})
		}
		_ = rows.Close()
	}

	return allCategory
}

func (cm *CategoryModel) getCategoryTree(parentID int) (error, []CategoryTree) {
	var categoryTree []CategoryTree
	categoryList := cm.QueryTargetParentTree(parentID)

	for _, v := range categoryList {
		_, trees := cm.getCategoryTree(v.CategoryID)
		v.Children = trees
		categoryTree = append(categoryTree, v)
	}
	return nil, categoryTree
}
