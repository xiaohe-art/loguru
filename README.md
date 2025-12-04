# loguru (Go version)

一个 **高性能、轻量、Loguru 风格** 的 Go 日志库封装，基于 Zerolog +
Lumberjack，实现了：

-   🌈 **开发环境彩色 Console 输出**
-   📁 **生产环境文件切割输出（lumberjack）**
-   🎯 **简洁的 loguru-style API：`Info(v)`，不用写格式化语句**
-   🧩 支持 **任意类型**：`string`, `int`, `struct`, `map`, `slice`,
    `interface{}`
-   📦 **Console + File 双通道输出**
-   🧱 自动添加时间戳，生产环境自动记录 Caller
-   ⚡ 零反射、极高性能（Zerolog 底层支持）

## ✨ 安装

``` bash
go get github.com/xiaohe-art/loguru
```

## 🚀 快速开始

``` go
package main

import "github.com/xiaohe-art/loguru/logger"

func main() {
    logger.Init(true, "logs/app.log") // 开发模式 + 写入文件

    logger.Info("hello world")
    logger.Warn("warning:", 123)
    logger.Error("something wrong")

    // 任意数据结构
    logger.Info(map[string]any{
        "user": "xiaohe",
        "age":  18,
    })
}
```

# 🕹 API 说明

所有 API 都支持 **任意类型** 的数据：

``` go
func Info(v any, args ...any)
func Debug(v any, args ...any)
func Warn(v any, args ...any)
func Error(v any, args ...any)
func Fatal(v any, args ...any)
```

## 示例

### 多参数自动拼接

``` go
logger.Info("login", "user:", "xiaohe", "age:", 18)
```

### 复杂对象

``` go
type User struct {
    Name string
    Age  int
}

logger.Info(User{"xiaohe", 20})
```

### 使用 map / slice

``` go
logger.Info(map[string]any{
    "event": "signup",
    "ip": "127.0.0.1",
})
```

# 📁 日志文件输出（生产环境）

``` go
logger.Init(false, "/var/log/myapp/app.log")
```

# 🌈 开发环境：彩色控制台输出

``` go
logger.Init(true, "app.log")
```

# 📦 目录结构建议

    loguru/
        logger/
            logger.go
        go.mod
        README.md

# 🏷 发布 Tag

``` bash
git tag v1.0.0
git push origin v1.0.0
```
