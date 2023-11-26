package main

import (
    "os"
    "fmt"
    "plugin" // 1. 导入标准库
)

// https://pkg.go.dev/plugin

// 2. 为插件定义接口
type Greeter interface {
    Greet(log string)
}

type HotFixPrint func(string)

func main() {

    // 3. 查找并实例化插件
    plug, err := plugin.Open("../HotFix/greet.so")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // 4. 找到插件导出的接口实例，其实这个不是必须的
    symGreeter, err := plug.Lookup("Greeter")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // 5. 类型转换
    var greeter Greeter
    greeter, ok := symGreeter.(Greeter)
    if !ok {
        fmt.Println(err)
        os.Exit(1)
    }

    // 6. 调用方法
    greeter.Greet("from Server main")

    symPrint, err := plug.Lookup("HotFixPrint")
    if err != nil {
        fmt.Println(err)
        return
    }

    hPrint := symPrint.(func(string))
    hPrint("From Server log")
}