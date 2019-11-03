package main

import (
    "errors"
    "fmt"

    "ccovers/http_server/logic"
    "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("服务器启动")

    init_err = global.InitLogicDb()
    defer global.CloseLogicDb()
    if init_err != nil {
        goto EXIT_PROCESS
    }
    fmt.Println("logic初始化完成")

    // 注册一个默认的路由器
    engin = gin.Default()
    regist_router(engine)
    engine.Run(":8080")

EXIT_PROCESS:
    fmt.Println("初始化异常，退出程序")
    fmt.Println(init_err.Error())
    os.Exit(-1)
}

func regist_router(engine *gin.Engine) {
    if engine == nil {
        return errors.New("无效的GIN实例")
    }

    router := engine.Group("/v1")
    {
        router.POST("/get_user_info", logic.GetUserInfo)
    }
}
