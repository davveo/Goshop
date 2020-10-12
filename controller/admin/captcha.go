package admin

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type MyCaptcha struct {
}

var Store = base64Captcha.DefaultMemStore

func NewDriver() *base64Captcha.DriverString {
	driver := new(base64Captcha.DriverString)
	driver.Height = 44
	driver.Width = 120
	driver.NoiseCount = 10
	driver.ShowLineOptions = base64Captcha.OptionShowSineLine |
		base64Captcha.OptionShowHollowLine
	driver.Length = 4
	driver.Source = "1234567890"
	driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

func NewCaptcha(context *gin.Context) {
	var driver = NewDriver().ConvertFonts()
	uuid := context.Param("uuid")
	scene := strings.ToUpper(context.Param("scene"))

	uniqueKey := fmt.Sprintf("%s_%s", scene, uuid)

	c := base64Captcha.NewCaptcha(driver, Store)
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, _ := c.Driver.DrawCaptcha(content)
	c.Store.Set(uniqueKey, answer)
	_, _ = item.WriteTo(context.Writer)
}
