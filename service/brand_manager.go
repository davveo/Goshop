package service

import "Goshop/model"

type BrandManager struct {

}

func (bm *BrandManager) list(page, pageSize int, name string) []map[string]interface{} {

}

func (bm *BrandManager) add(brand model.BrandModel) {

}

func (bm *BrandManager) edit(brand model.BrandModel, id int) {

}


func (bm *BrandManager) delete(ids []int) {

}


func (bm *BrandManager) getModel(id int) {

}


func (bm *BrandManager) getBrandsByCategory(categoryId int) []model.BrandModel {
	return nil
}


func (bm *BrandManager) getCatBrand(categoryId int) {

}

func (bm *BrandManager)getAllBrands()[]model.BrandModel  {
	return nil
}
