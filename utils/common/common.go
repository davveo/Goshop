package common

import (
	"Goshop/global/variable"
	"crypto/sha1"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func ParseFromQuery(ctx *gin.Context) map[string]string {
	rs := make(map[string]string)
	for key, value := range ctx.Request.URL.Query() {
		rs[key] = value[0]
	}
	if _, ok := rs["page_no"]; !ok {
		rs["page_no"] = string(variable.PageNum) // 默认页码
	}
	if _, ok := rs["page_size"]; !ok {
		rs["page_size"] = string(variable.PageSize) // 默认条数
	}
	return rs
}

func ParseFromBody(ctx *gin.Context) map[string]interface{} {
	rs := make(map[string]interface{})
	if ctx.Request.PostForm == nil {
		return rs
	}

	for key, value := range ctx.Request.PostForm {
		if len(value) == 1 {
			rs[key] = value[0]
		} else {
			rs[key] = value
		}

	}
	return rs
}
