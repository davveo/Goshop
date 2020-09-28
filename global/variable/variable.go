package variable

import (
	"log"
	"Goshop/global/my_errors"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
)

var (
	AccessTokenExpireTime       int64         = 1200
	RefreshTokenExpireTime      int64         = 1800
	AccessTokenCacheExpireTime  time.Duration = 1260
	RefreshTokenCacheExpireTime time.Duration = 1860
	BasePath                    string        // 定义项目的根目录
	EventDestroyPrefix          = "Destroy_"  //  程序退出时需要销毁的事件前缀
	//上传文件保存路径
	UploadFileField    = "files"                  // post上传文件时，表单的键名
	UploadFileSavePath = "/storage/app/uploaded/" // 该路径与 BasePath 进行拼接使用

	//日志存储路径
	ZapLog *zap.Logger //  全局日志句柄

	//websocket
	WebsocketHub              interface{}
	WebsocketHandshakeSuccess = "Websocket Handshake+OnOpen Success"
	WebsocketServerPingMsg    = "Server->Ping->Client"
	//  用户自行定义其他全局变量 ↓

	AccessTokenPrefix   = "{ACCESS_TOKEN}_ADMIN_%s_%s"
	RefreshTokenPrefix  = "{REFRESH_TOKEN}_ADMIN_%s_%s"
	AdminDisabledPrefix = "{ADMIN_DISABLED}_ADMIN_%s"
	SettingsPrefix      = "{SETTING}_%s"
)

func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = path
		}
	} else {
		log.Fatal(my_errors.ErrorsBasePath)
	}
}
