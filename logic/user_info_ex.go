package logic

import (
    "fmt"
    "strconv"

    "ccovers/http_server/lib"

    "github.com/gin-gonic/gin"
)

func GetUserInfoEx(c *gin.Context) {
    res := GetUserInfoRes{}
    var err error

    for ok := true; ok; ok = false {
        var userId int64
        idStr := c.Param("userId")
        userId, err = strconv.ParseInt(idStr, 10, 64)
        if err != nil {
            fmt.Printf("解析错误: %s, %s\n", idStr, err.Error())
            break
        }

        err = getUserInfo(&GetUserInfoReq{
            UserId: userId,
        }, &res)
        if err != nil {
            fmt.Printf("查询错误: %s\n", err.Error())
            break
        }
        lib.ServerResponse(c, nil, &res)
        return
    }
    lib.ServerResponse(c, err, nil)
}
