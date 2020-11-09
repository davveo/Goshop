package consts

// GoodsChangeMsg operationType
const (
	OperationAddOperation = iota + 1

	/**
	 * 修改
	 */
	OperationUpdateOperation

	/**
	 * 删除
	 */
	OperationDelOperation

	/**
	 * 下架
	 */
	OperationUpOperation

	/**
	 * 下架
	 */
	OperationUnderOperation

	/**
	 * 还原
	 */
	OperationRevertOperation

	/**
	 * 放入回收站
	 */
	OperationInrecycleOperation

	/**
	 * 商品成功审核
	 */
	OperationGoodsVerifySuccess

	/**
	 * 商品失败审核
	 */
	OperationGoodsVerifyFail

	/**
	 * 商品失败审核
	 */
	OperationGoodsPriorityChange
)
