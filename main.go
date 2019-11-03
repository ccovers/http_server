package main

import (
    "errors"
    "fmt"
    "os"

    "ccovers/http_server/global"
    "ccovers/http_server/logic"
    "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("服务器启动")

    init_err := global.InitLogicDb()
    defer global.CloseLogicDb()
    if init_err != nil {
        goto EXIT_PROCESS
    }
    fmt.Println("logic初始化完成")

    startServer()

EXIT_PROCESS:
    fmt.Println("初始化异常，退出程序")
    fmt.Println(init_err.Error())
    os.Exit(-1)
}

func startServer() {
    gin.SetMode(gin.ReleaseMode)

    // 注册一个默认的路由器
    engine := gin.Default()
    registRouter(engine)
    engine.Run(":8080")
}

func registRouter(engine *gin.Engine) error {
    if engine == nil {
        return errors.New("无效的GIN实例")
    }

    router := engine.Group("/v1")
    {
        router.POST("/get_user_info", logic.GetUserInfo)
        router.GET("/get_user_info_ex", logic.GetUserInfoEx)
    }
    return nil
}
