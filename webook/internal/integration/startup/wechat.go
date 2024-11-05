package startup

import (
	"github.com/ischeng28/basic-go/webook/internal/service/oauth2/wechat"
	"github.com/ischeng28/basic-go/webook/pkg/logger"
)

func InitWechatService(l logger.LoggerV1) wechat.Service {
	return wechat.NewService("", "", l)
}
