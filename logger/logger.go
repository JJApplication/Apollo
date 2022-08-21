/*
Project: Apollo logger.go
Created: 2021/11/21 by Landers
*/

package logger

// 日志记录器 zap
import (
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var LoggerSugar *zap.SugaredLogger

func InitLogger() error {
	logger, err := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableStacktrace: configStack(),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: configEncoding(),
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:          "Time",
			LevelKey:         "Level",
			NameKey:          "Name",
			CallerKey:        configCaller(),
			MessageKey:       "Message",
			FunctionKey:      configFunction(),
			StacktraceKey:    "Stack",
			EncodeName:       nil,
			ConsoleSeparator: "",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      zapcore.CapitalLevelEncoder,
			EncodeTime:       zapcore.TimeEncoderOfLayout(utils.TimeForLogger),
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
		},
		OutputPaths:      configLog(),
		ErrorOutputPaths: configLog(),
		InitialFields:    map[string]interface{}{"Name": LoggerPrefix},
	}.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		logger.Error(LoggerInitFailed)
		return err
	}

	Logger = logger
	LoggerSugar = logger.Sugar()
	Logger.Info(LoggerInitSuccess)
	defer logger.Sync()

	return nil
}

func configEncoding() string {
	switch config.ApolloConf.Log.Encoding {
	case "json", "JSON":
		return "json"
	case "console", "Console":
		return "console"
	default:
		return "json"
	}
}

func configFunction() string {
	switch config.ApolloConf.Log.EnableFunction {
	case "yes", "YES", "Yes":
		return "Function"
	case "no", "NO", "No":
		return ""
	default:
		return ""
	}
}

func configLog() []string {
	switch config.ApolloConf.Log.EnableLog {
	case "yes", "YES", "Yes":
		if config.ApolloConf.Log.LogFile != "" {
			return []string{"stdout", config.ApolloConf.Log.LogFile}
		}
		return []string{"stdout"}
	case "no", "NO", "No":
		return []string{}
	default:
		return []string{"stdout"}
	}
}

func configStack() bool {
	switch config.ApolloConf.Log.EnableStack {
	case "yes", "YES", "Yes":
		return false
	case "no", "NO", "No":
		return true
	default:
		return true
	}
}

func configCaller() string {
	switch config.ApolloConf.Log.EnableCaller {
	case "yes", "YES", "Yes":
		return "Caller"
	case "no", "NO", "No":
		return ""
	default:
		return ""
	}
}
