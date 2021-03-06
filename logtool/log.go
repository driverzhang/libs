package logtool

import (
	"time"
	"go.uber.org/zap/zapcore"
	"os"
	"go.uber.org/zap"
	"path/filepath"
)

var Zap *zap.Logger

func InitZapLogger(lpName string, isDebug bool) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var cfg zap.Config
	if isDebug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.Development = true
		cfg.Encoding = "console"
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stderr"}
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		cfg.Development = false
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		cfg.Sampling = &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		}
		cfg.Encoding = "json"
		cfg.OutputPaths = []string{"stdout", filepath.Join(path, lpName)}
		cfg.ErrorOutputPaths = []string{"stderr"}
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	}
	cfg.EncoderConfig.EncodeTime = timeEncoder
	Zap, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
}
