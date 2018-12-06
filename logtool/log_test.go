package logtool

import (
	"testing"
	"go.uber.org/zap"
	"errors"
)

func TestInitLogger(t *testing.T) {
	logTool, err := InitLogger("logtool.log", true)
	if err != nil {
		panic(err)
	}
	d := []string{" info", " error", " debug", " fatal",}
	logTool.Info("info:", zap.Error(errors.New("123123")))
	logTool.Error("info:", zap.String("s", d[1]))
	logTool.Debug("info:", zap.String("s", d[2]))
	logTool.Fatal("info:", zap.String("s", d[3]))
}
