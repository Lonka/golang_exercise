package logger

import(
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
	"os"
)

func main(){
	
	hook := lumberjack.Logger{
		Filename:"./logs/spikeProxy.log",// 日誌檔案路徑
		MaxSize:128,// 每個日誌檔案儲存的最大尺寸 單位：M
		MaxBackups:30, // 日誌檔案最多儲存多少個備份
		MaxAge:7,// 檔案最多儲存多少天
		Compress:true,// 是否壓縮
	}

	atom := zap.NewAtomicLevelAt(zap.InfoLevel)
	core := zapcore.NewCore(
						zapcore.NewJSONEncoder(NewDevelopmentEncoderConfig()),// 編碼器配置
						zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(&hook)),// 列印到控制檯和檔案
						atom,// 日誌級別
					)

	// 開啟開發模式，堆疊跟蹤				
	caller := zap.AddCaller()
	// 開啟檔案及行號
	development := zap.Development()
	// 設定初始化欄位
	filed := zap.Fields(zap.String("serviceName","serviceName"))
	// 構造日誌
	logger := zap.New(core,caller,development,filed)


	logger.Info("log 初始化成功")

	logger.Info("无法获取网址",
			zap.String("url", "http://www.baidu.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
	)

	logger.Error("error",zap.String("key","value"))


}

func NewDevelopmentEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // ISO8601 UTC 時間格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}



func NewProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,// 小寫編碼器
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
