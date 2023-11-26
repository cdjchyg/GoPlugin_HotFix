package main

import (
    "fmt"
    "plugin"
    "GoHotFix/Server/player"
    "GoHotFix/Struct"
)

func loadDispatchMsg(path string, funcName string) func(interface{})interface{} {
    plug, err := plugin.Open(path)
    if err != nil {
        return nil
    }

    symDispatchMsg, err := plug.Lookup(funcName)
    if err != nil {
        return nil
    }

    dispatchMsg := symDispatchMsg.(func(interface{})interface{})
    return dispatchMsg
}

func main() {
    dispatchMsg := loadDispatchMsg("../HotFix/hotfix.so.1", "DispatchMsg")

    pm := player.NewPlayerMgr()
    dispatchMsg(&Struct.SetPlayerMgrReq{PM: pm})

    // 模拟客户端过来的消息
    dispatchMsg(&Struct.CreatePlayerReq{PID: 1})
    dispatchMsg(&Struct.UpdateNameReq{PID:1, Name:"name1"})
    ack := dispatchMsg(&Struct.GetPlayerNameReq{PID:1}).(*Struct.GetPlayerNameAck)
    fmt.Println(ack.Name)

    fmt.Println("reload ===================")

    // reload for hot fix
    dispatchMsg = loadDispatchMsg("../HotFix/hotfix.so.2", "DispatchMsg")
    dispatchMsg(&Struct.SetPlayerMgrReq{PM: pm})
    ack = dispatchMsg(&Struct.GetPlayerNameReq{PID:1}).(*Struct.GetPlayerNameAck)
    fmt.Println(ack.Name)
}

// // https://pkg.go.dev/plugin
    //// Plugins are currently supported only on Linux, FreeBSD, and macOS, making them unsuitable for applications intended to be portable.

// // 2. 为插件定义接口
// type Greeter interface {
//     Greet(log string)
// }

// type HotFixPrint func(string)

// func main() {

//     // 3. 查找并实例化插件
//     plug, err := plugin.Open("../HotFix/greet.so")
//     if err != nil {
//         fmt.Println(err)
//         os.Exit(1)
//     }

//     // 4. 找到插件导出的接口实例，其实这个不是必须的
//     symGreeter, err := plug.Lookup("Greeter")
//     if err != nil {
//         fmt.Println(err)
//         os.Exit(1)
//     }

//     // 5. 类型转换
//     var greeter Greeter
//     greeter, ok := symGreeter.(Greeter)
//     if !ok {
//         fmt.Println(err)
//         os.Exit(1)
//     }

//     // 6. 调用方法
//     greeter.Greet("from Server main")

//     symPrint, err := plug.Lookup("HotFixPrint")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }

//     hPrint := symPrint.(func(string))
//     hPrint("From Server log")
// }