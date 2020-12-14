package consts

var (
	FlowMap = map[string]interface{}{
		// 取消订单 （订单付款之后，确认收货之前）
		"ORDER_CANCEL": map[string]interface{}{
			"APPLY":           []string{"SELLER_AUDIT"},  // 申请状态，商家可以审核
			"PASS":            []string{"STOCK_IN"},      // 审核通过状态，商家可以入库
			"STOCK_IN":        []string{"SELLER_REFUND"}, // 入库状态，商家可以退款
			"WAIT_FOR_MANUAL": []string{"ADMIN_REFUND"},  // 等待人工处理状态，管理员可以退款
			"ERROR_EXCEPTION": []string{"CLOSE"},         // 异常状态 可以关闭售后服务
			"CLOSED":          []string{},                // 关闭状态不能进行任何操作
			"REFUSE":          []string{},                // 审核未通过状态不能进行任何操作
			"COMPLETED":       []string{},                // 完成状态不能进行任何操作
		},
		// 退货流程
		"RETURN_GOODS": map[string]interface{}{
			"APPLY":           []string{"SELLER_AUDIT"},        // 申请状态，商家可以审核
			"PASS":            []string{"FILL_LOGISTICS_INFO"}, // 审核通过状态，买家可以填写物流信息
			"FULL_COURIER":    []string{"STOCK_IN"},            // 物流完善状态可以确认入库
			"STOCK_IN":        []string{"SELLER_REFUND"},       // 入库状态，商家可以退款
			"WAIT_FOR_MANUAL": []string{"ADMIN_REFUND"},        // 等待人工处理状态，管理员可以退款
			"ERROR_EXCEPTION": []string{"CLOSE"},               // 异常状态 可以关闭售后服务
			"CLOSED":          []string{},                      // 关闭状态不能进行任何操作
			"REFUSE":          []string{},                      // 审核未通过状态不能进行任何操作
			"COMPLETED":       []string{},                      // 完成状态不能进行任何操作
		},
		// 换货流程
		"CHANGE_GOODS": map[string]interface{}{
			"APPLY":           []string{"SELLER_AUDIT"},              //  申请状态，商家可以审核
			"PASS":            []string{"FILL_LOGISTICS_INFO"},       // 审核通过状态，买家可以填写物流信息
			"FULL_COURIER":    []string{"STOCK_IN"},                  // 物流完善状态可以确认入库
			"ERROR_EXCEPTION": []string{"CREATE_NEW_ORDER", "CLOSE"}, // 异常状态 商家可以手动创建新订单也可以关闭售后服务单
			"CLOSED":          []string{},                            // 关闭状态不能进行任何操作
			"REFUSE":          []string{},                            // 审核未通过状态不能进行任何操作
			"COMPLETED":       []string{},                            // 完成状态不能进行任何操作
		},
		// 补发商品流程
		"SUPPLY_AGAIN_GOODS": map[string]interface{}{
			"APPLY":           []string{"SELLER_AUDIT"},              // 申请状态，商家可以审核
			"ERROR_EXCEPTION": []string{"CREATE_NEW_ORDER", "CLOSE"}, // 异常状态 商家可以手动创建新订单也可以关闭售后服务单
			"CLOSED":          []string{},                            // 关闭状态不能进行任何操作
			"REFUSE":          []string{},                            // 审核未通过状态不能进行任何操作
			"COMPLETED":       []string{},                            // 完成状态不能进行任何操作
		},

		// 款到发货流程
		"ONLINE": map[string]interface{}{
			"NEW":           []string{"CONFIRM", "CANCEL"},      // 新订单，可以确认，可以取消
			"CONFIRM":       []string{"PAY", "CANCEL"},          // 确认的订单，可以支付，可以取消
			"PAID_OFF":      []string{"SHIP", "SERVICE_CANCEL"}, // 已经支付，可以发货，可以取消订单
			"SHIPPED":       []string{"ROG", "SERVICE_CANCEL"},  // 发货的订单，可以确认收货，可以取消订单
			"ROG":           []string{"COMPLETE", "COMMENT"},    // 收货的订单，可以完成,可以评论(此时需要校验评论状态是否完成)
			"AFTER_SERVICE": []string{"COMPLETE"},               //售后的订单，可以完成
			"CANCELLED":     []string{},                         // 取消的的订单不能有任何操作
			"INTODB_ERROR":  []string{},                         // 异常的订单不能有任何操作
			"COMPLETE":      []string{},                         //完成的订单不能有任何操作
		},
		//货到付款流程
		"COD": map[string]interface{}{
			"NEW":           []string{"CONFIRM", "CANCEL"}, // 新订单，可以确认，可以取消
			"CONFIRM":       []string{"SHIP", "CANCEL"},    // 确认的订单，可以发货,可取消
			"PAID_OFF":      []string{"COMPLETE"},          // 收货的订单，可以完成
			"SHIPPED":       []string{"ROG"},               // 发货的订单，可以确认收货
			"ROG":           []string{"PAY"},               // 收货的订单，可以支付
			"AFTER_SERVICE": []string{"COMPLETE"},          //售后的订单，可以完成
			"CANCELLED":     []string{},                    // 取消的的订单不能有任何操作
			"INTODB_ERROR":  []string{},                    // 异常的订单不能有任何操作
			"COMPLETE":      []string{},                    //完成的订单不能有任何操作
		},
		//拼团的订单流程
		"PINTUAN": map[string]interface{}{
			"NEW":           []string{"CONFIRM", "CANCEL"}, // 新订单，可以确认，可以取消
			"CONFIRM":       []string{"PAY", "CANCEL"},     // 确认的订单，可以支付，可以取消
			"PAID_OFF":      []string{"SERVICE_CANCEL"},    // 已经支付，可以发货，可以取消订单
			"FORMED":        []string{"SHIP"},              // 已经成团的，可以发货
			"SHIPPED":       []string{"ROG"},               // 发货的订单，可以确认收货
			"ROG":           []string{"COMPLETE"},          // 收货的订单，可以完成
			"AFTER_SERVICE": []string{"COMPLETE"},          //售后的订单，可以完成
			"CANCELLED":     []string{},                    // 取消的的订单不能有任何操作
			"INTODB_ERROR":  []string{},                    // 异常的订单不能有任何操作
			"COMPLETE":      []string{},                    //完成的订单不能有任何操作
		},
	}
)
