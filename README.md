# github.com/thehappymouse/go-utils

#### 项目介绍
集工作、学习中的的通用工具方法


#### 使用说明

> 方式一
```
# go get -u github.com/thehappymouse/go-util

或者使用 mod 

require (
	github.com/magiconair/properties v1.8.1
	github.com/rs/zerolog v1.14.3
	github.com/thehappymouse/go-utils v0.0.2
)

# 代码文件
import "github.com/thehappymouse/go-utils"
# 命令行打印
utils.ZeroConsoleLog()
```
> 方试二
```go.mod
# cat go.mod
module modulename

go 1.12

require (
	github.com/magiconair/properties v1.8.1
	github.com/rs/zerolog v1.14.3
	github.com/thehappymouse/go-utils v0.0.2
)

```

#### 内容简述
* 36进制
* zero.log 命令行输出格式化
* http 下载文件
