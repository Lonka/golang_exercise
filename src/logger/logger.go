package logger

import(
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var MainLogger *zap.Logger

func init(){
	MainLogger = NewLogger("./logs/main.log",zapcore.InfoLevel,128,30,7,true,"Main")
}
