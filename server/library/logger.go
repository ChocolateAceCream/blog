// TODO: add file rotate using  github.com/lestrrat-go/file-rotatelogs
package library

import (
	"fmt"
	"os"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logTmFmtWithMS = "2006-01-02 15:04:05.000"
)

func GetEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  global.CONFIG.Zap.StacktraceKey,
		EncodeTime:     GetCustomTimeEncoder, // 自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   GetCustomCallerEncoder, // caller trace encoder
	}
	GetEncodeLevel(&config)
	return config
}

// 自定义文件：行号输出项
func GetCustomCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

// Set logger encode level based on config
func GetEncodeLevel(c *zapcore.EncoderConfig) {
	switch global.CONFIG.Zap.EncodeLevel {
	case "LowercaseLevelEncoder":
		c.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		c.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		c.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		c.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
}

// Set CustomTimeEncoder
func GetCustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format(global.CONFIG.Zap.Prefix + " 2006/01/02 - 15:04:05.000 "))
}

func GetOutputPath() string {
	if ok, _ := utils.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}
	return fmt.Sprintf("./%s/%s", global.CONFIG.Zap.Director, global.CONFIG.Zap.FileName)
}

func LoggerInit() *zap.Logger {
	encoderConfig := GetEncoderConfig()

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.InfoLevel)

	config := zap.Config{
		Level:       atom,  // log level
		Development: false, // 开发模式，堆栈跟踪

		// DisableStacktrace completely disables automatic stacktrace capturing. By
		// default, stacktraces are captured for WarnLevel and above logs in
		// development and ErrorLevel and above in production.
		DisableStacktrace: true,
		Encoding:          global.CONFIG.Zap.Format, // console or json
		EncoderConfig:     encoderConfig,
		InitialFields:     map[string]interface{}{"serviceName": "wsl"}, // add new key-value pair
		OutputPaths:       []string{"stdout", GetOutputPath()},          //  stdout and customized destination
		ErrorOutputPaths:  []string{"stderr"},
	}

	// build log
	ZapLog_V1, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed log initialization : %v", err))
	}

	// testing
	ZapLog_V1.Warn("warn: log initialization successful")
	ZapLog_V1.Info("info: log initialization successful")
	ZapLog_V1.Error("err: log initialization successful")

	// print out error
	// if _, err := strconv.Atoi("a123"); err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	ZapLog_V1.Error(fmt.Sprintf("err: %v", err))
	// }
	return ZapLog_V1
}
