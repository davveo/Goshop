package model

import (
	"Goshop/global/consts"
	"Goshop/utils/rabbitmq"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"errors"
	"fmt"
	"log"
	"strings"
)

func CreateCategoryFactory(sqlType string) *CategoryModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	mq := rabbitmq.GetRabbitmq()
	if mq == nil {
		log.Fatal("goodsModel mq初始化失败")
	}
	amqpTemplate, err := mq.Producer("category")
	if err != nil {
		log.Fatal("categoryModel producer初始化失败")
	}

	if dbDriver != nil {
		return &CategoryModel{
			BaseModel:    dbDriver,
			amqpTemplate: amqpTemplate,
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
	amqpTemplate *rabbitmq.Producer
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

func (cm *CategoryModel) getModel(parentID string) (map[string]interface{}, error) {
	sqlString := "select * from es_category where parent_id = ?"
	rows := cm.QuerySql(sqlString, parentID)
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

func (cm *CategoryModel) Add(params map[string]interface{}) (map[string]interface{}, error) {
	var (
		categoryId int64
		parent     map[string]interface{}
	)

	name := params["name"].(string)
	image := params["image"].(string)
	isShow := params["is_show"].(string)
	advImage := params["advImage"].(string)
	parentId := params["parent_id"].(string)
	advImageLink := params["advImageLink"].(string)
	categoryOrder := params["category_order"].(string)

	if name == "" {
		return nil, errors.New("分类名不能为空")
	}

	sqlString := "select * from es_category where name = ? "
	rows := cm.QuerySql(sqlString, name)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, err
	}

	if len(tableData) != 0 {
		return nil, errors.New("分类名已经存在")
	}
	// 非顶级分类
	if parentId != "" && parentId != "0" {
		parent, _ = cm.getModel(parentId)
		if parent == nil {
			return nil, errors.New("父分类不存在")
		}
		catPath := parent["category_path"].(string)
		str := strings.Split(strings.Replace(
			catPath, "|", ",", -1), ",")

		if len(str) >= 4 {
			return nil, errors.New("最多为三级分类，添加失败")
		}
	}

	sqlString = "insert into `es_category` " +
		"(`name`, `parent_id`, `is_show`, `category_order`, `image`, `adv_image`, `adv_image_link`) " +
		"values (?,?,?,?,?,?,?)"

	if categoryId = cm.LastInsertId(sqlString, name, parentId, isShow,
		categoryOrder, image, advImage, advImageLink); categoryId == -1 {
		return nil, errors.New("插入分类失败")
	} else {
		params["category_id"] = categoryId
	}
	// 判断是否是顶级类似别，如果parentid为空或为0则为顶级类似别
	// 注意末尾都要加|，以防止查询子孙时出错
	// 不是顶级类别，有父
	if parent != nil {
		params["category_path"] = parent["category_path"].(string) + string(categoryId) + "|"
	} else { // 顶级分类
		params["category_path"] = "0|" + string(categoryId) + "|"
	}

	sqlString = "update es_category set  category_path=? where  category_id=?"
	if affected := cm.ExecuteSql(sqlString, params["category_path"].(string), categoryId); affected == -1 {
		return nil, errors.New("更新分类失败")
	}
	rds.Remove(fmt.Sprintf("%s_%s", consts.GOODS_CAT, "ALL"))

	// 发送消息变化消息
	categoryChangeMsg := rabbitmq.BuildMsg(map[string]interface{}{
		"category_id":    categoryId,
		"operation_type": consts.OperationAddOperation,
	})
	err = cm.amqpTemplate.Publish(consts.ExchangeGoodsCategoryChange,
		consts.ExchangeGoodsCategoryChange+"_ROUTING", categoryChangeMsg)
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	return params, nil
}
