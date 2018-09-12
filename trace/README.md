## Opentracing

### 初始化

```go
//初始化
func TraceInit() {
	//配置
	c = trace.TraceingInit("127.0.0.1:6831","project_name")
}
//关闭
func Close() {
	c.Close()
}
```

### 使用

```go

//opentracing从context获取,写入context，适用RPC
	ctx, span, _ := utils.TraceIntoContext(ctx, "test")
	defer span.Finish()

//opentracing从header获取,写入context,适用获取http
	ctx, span, _ := utils.TraceFromHeader(ctx, "test",header)
	defer span.Finish()
	
//opentracing从context获取,写入http,适用将调用http
	ctx, span, _ := utils.TraceToHeader(ctx, "test",header)
	defer span.Finish()
```