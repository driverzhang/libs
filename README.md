# 实用工具库
### perf 　开启http pprof进行系统监控
```base
    StartPprof([]string{"127.0.0.1:8077"})
```
### quit  优雅的退出go程序 
```base
    quit.QuitSignal(func() {
		fmt.Println("退出程序")
	})
```
### zap 日志库相关
```base
InitZapLogger("logtool.log", true)
	d := []string{" info", " error", " debug", " fatal",}
	Zap.Info("info:", zap.Error(errors.New("123123")))
	Zap.Error("info:", zap.String("s", d[1]))
	Zap.Debug("info:", zap.String("s", d[2]))
	Zap.Fatal("info:", zap.String("s", d[3]))
```