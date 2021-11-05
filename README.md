## 目的

演示如何使用`Go Module`创建项目
`Go Module`是从 1.11 版引入，是`GOPATH`的替代品，集成版本控制和软件包分发支持
`GOPATH`已弃用

## 使用过程

### 创建项目

```bash
mkdir goapp1 && cd $_
go mod init goapp1
```

`go mod` 将创建`go.mod`

### 第三方依赖

```go
package main

import (
    "fmt"
    controllers "goapp1/handlers"

    jsoniter "github.com/json-iterator/go"
)

type AppInfo struct {
    Name string
}

func main() {
    info := AppInfo{
        Name: "GoApp",
    }
    jsonString, _ := jsoniter.Marshal(&info)

    fmt.Println(string(jsonString))

}
```

执行`go get github.com/json-iterator/go`
将创建`go.sum`，同时更新`go.mod`

### 本地依赖

```bash
mkdir handlers && touch showInfo.go
```

在 showInfo.go 中，创建：

```go
package controllers

import "fmt"

func ShowInfo() {
    fmt.Println("controller modeule info")
}
```

使用本地依赖

```go
package main

import (
    "fmt"
    controllers "goapp1/handlers"

    jsoniter "github.com/json-iterator/go"
)

type AppInfo struct {
    Name string
}

func main() {
    info := AppInfo{
        Name: "GoApp",
    }
    jsonString, _ := jsoniter.Marshal(&info)

    fmt.Println(string(jsonString))
    controllers.ShowInfo()
}
```
