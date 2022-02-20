/*
Project: dirichlet logger.go
Created: 2021/11/21 by Landers
*/

package logger

// 日志记录器 zap
import (
	"github.com/landers1037/dirichlet/config"
	"github.com/landers1037/dirichlet/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

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
		InitialFields:    map[string]interface{}{"Logger": LoggerPrefix},
	}.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		logger.Error(LoggerInitFailed)
		return err
	}

	Logger = logger
	Logger.Info(LoggerInitSuccess)
	defer logger.Sync()

	return nil
}

func configEncoding() string {
	switch config.DirichletConf.Log.Encoding {
	case "json", "JSON":
		return "json"
	case "console", "Console":
		return "console"
	default:
		return "json"
	}
}

func configFunction() string {
	switch config.DirichletConf.Log.EnableFunction {
	case "yes", "YES", "Yes":
		return "Function"
	case "no", "NO", "No":
		return ""
	default:
		return ""
	}
}

func configLog() []string {
	switch config.DirichletConf.Log.EnableLog {
	case "yes", "YES", "Yes":
		if config.DirichletConf.Log.LogFile != "" {
			return []string{"stderr", config.DirichletConf.Log.LogFile}
		}
		return []string{"stderr"}
	case "no", "NO", "No":
		return []string{}
	default:
		return []string{"stderr"}
	}
}

func configStack() bool {
	switch config.DirichletConf.Log.EnableStack {
	case "yes", "YES", "Yes":
		return false
	case "no", "NO", "No":
		return true
	default:
		return true
	}
}

func configCaller() string {
	switch config.DirichletConf.Log.EnableCaller {
	case "yes", "YES", "Yes":
		return "Caller"
	case "no", "NO", "No":
		return ""
	default:
		return ""
	}
}
