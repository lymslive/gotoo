# 一些 go 语言的小工具辅助库
`+` `go`

自己在使用 go 语言实践中提取的一些实用小库。

## assert

便于在写单元测试中写 assert 断言。用法如：

```go
import "github.com/lymslive/gotoo/assert"

assert.BeginTest(t)
assert.True(val, msg)
assert.Equal(got, expect, msg)
assert.EndTest()
```
