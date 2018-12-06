package logtool

import (
	"testing"
	"go.uber.org/zap"
	"errors"
)

func TestInitLogger(t *testing.T) {
InitZapLogger("logtool.log", true)
	d := []string{" info", " error", " debug", " fatal",}
	Zap.Info("info:", zap.Error(errors.New("123123")))
	Zap.Error("info:", zap.String("s", d[1]))
	Zap.Debug("info:", zap.String("s", d[2]))
	Zap.Fatal("info:", zap.String("s", d[3]))
}
