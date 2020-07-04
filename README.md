# summer-demo

```
    // 注册 controller
	summer.App.SetController("/hello", func(ctx context.Context, request *http.Request) string {
		return "hello"
	})

    // 绑定端口
	summer.App.SetPort("8080")

    // 程序启动
	summer.App.Run()
```