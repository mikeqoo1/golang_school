package file

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//LOG ...
var LOG *zap.Logger

//InitLog 初始化log
func InitLog() {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, //時間格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	// 設定log等級
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom, // log等級
		Development:      true,
		Encoding:         "json", // 輸出格式 console 或 json
		EncoderConfig:    encoderConfig,
		InitialFields:    map[string]interface{}{"Key值": "內容"}, // 初始化字段
		OutputPaths:      []string{"stdout", "test.log"},       // 輸出文件 stdout stderr
		ErrorOutputPaths: []string{"stderr", "error.log"},
	}

	// zap.NewDevelopment 格式化输出
	var err error
	LOG, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}
}
