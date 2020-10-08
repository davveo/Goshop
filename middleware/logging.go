package middleware

import (
	"Goshop/global/errno"
	"Goshop/utils/response"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now().UTC()
		path := ctx.Request.URL.Path

		// 过滤登录等接口
		//reg := regexp.MustCompile("(/v1/user|/login)")
		//if !reg.MatchString(path) {
		//	return
		//}

		//if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
		//	return
		//}

		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
		}

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		method := ctx.Request.Method
		ip := ctx.ClientIP()

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = blw

		// Continue.
		ctx.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		var resp response.Response
		if err := json.Unmarshal(blw.body.Bytes(), &resp); err != nil {
			log.Print(fmt.Sprintf(
				"response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes()))

			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = resp.Code
			message = resp.Message
		}
		log.Print(fmt.Sprintf(
			"%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message))
	}
}
