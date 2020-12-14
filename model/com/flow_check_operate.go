package com

import (
	"Goshop/global/consts"
)

func CheckOperate(flow, currentStatus, operate string) bool {
	statusInterface := getFlowMap(flow)
	if statusInterface == nil {
		return false
	}
	status := statusInterface.(map[string]interface{})
	if len(status) <= 0 {
		return false
	}

	allow := status[currentStatus].([]string)
	if len(allow) > 0 && InArr(operate, allow) {
		return true
	}
	return false
}

func InArr(operate string, arr []string) bool {
	for _, item := range arr {
		if item == operate {
			return true
		}
	}
	return false
}

func getFlowMap(flow string) interface{} {
	if status, ok := consts.FlowMap[flow]; ok {
		return status
	}
	return nil
}
