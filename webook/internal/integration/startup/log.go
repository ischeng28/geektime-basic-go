package startup

import "github.com/ischeng28/basic-go/webook/pkg/logger"

func InitLogger() logger.LoggerV1 {
	return logger.NewNopLogger()
}
