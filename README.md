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