package model

type OperateAllowable struct {
	disabled     int64
	marketEnable int64
}

func NewOperateAllowable(disabled int64, marketEnable int64) *OperateAllowable {
	return &OperateAllowable{disabled: disabled, marketEnable: marketEnable}
}

func (oa *OperateAllowable) getAllowUnder() bool {
	//上架并且没有删除的可以下架
	return oa.marketEnable == 1 && oa.disabled == 1
}

func (oa *OperateAllowable) getAllowRecycle() bool {
	//下架的商品才能放入回收站

	return oa.marketEnable == 0 && oa.disabled == 1
}

func (oa *OperateAllowable) getAllowRevert() bool {
	//下架的删除了的才能还原
	return oa.marketEnable == 0 && oa.disabled == 0
}

func (oa *OperateAllowable) getAllowDelete() bool {
	//下架的删除了的才能还原
	return oa.marketEnable == 0 && oa.disabled == 0
}

func (oa *OperateAllowable) getAllowMarket() bool {
	//下架未删除才能上架
	return oa.marketEnable == 0 && oa.disabled == 1
}
