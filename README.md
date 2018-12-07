# 实用工具库
```base
   go get github.com/hwholiday/libs 
```
### perf 　开启http pprof进行系统监控
```base
    import "github.com/hwholiday/libs/perf"
    
    perf.StartPprof([]string{"127.0.0.1:8077"})
```
### quit  优雅的退出go程序 
```base
    import "github.com/hwholiday/libs/quit"
    
    quit.QuitSignal(func() {
	    fmt.Println("退出程序")
    })
```
### zap 日志库相关
```base
    import "github.com/hwholiday/libs/logtool"
    
    logtool.InitZapLogger("logtool.log", true)
	d := []string{" info", " error", " debug", " fatal",}
	logtool.Zap.Info("info:", zap.Error(errors.New("123123")))
	logtool.Zap.Error("info:", zap.String("s", d[1]))
	logtool.Zap.Debug("info:", zap.String("s", d[2]))
	logtool.Zap.Fatal("info:", zap.String("s", d[3]))
```