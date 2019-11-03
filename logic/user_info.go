package logic

import (
    "fmt"

    "ccovers/http_server/lib"

    "github.com/gin-gonic/gin"
)

type GetUserInfoReq struct {
    UserId int64 `json:"userId" comment:"用户ID"`
}

type GetUserInfoRes struct {
    UserId  int64  `json:"userId" comment:"用户ID"`
    Name    string `json:"name" comment:"名称"`
    Subject string `json:"subject" comment:"科目"`
    Score   int    `json:"score" comment:"分数"`
}

func GetUserInfo(c *gin.Context) {
    req := GetUserInfoReq{}
    res := GetUserInfoRes{}

    for ok := true; ok; ok = false {
        err := lib.ReadRequestBodyJson(c, &req)
        if err != nil {
            fmt.Printf("解析错误: %s\n", err.Error())
            break
        }

        err = getUserInfo(&req)
        if err != nil {
            fmt.Println("查询错误: %s\n", err.Error())
            break
        }
        lib.ServerResponse(c, 0, &res)
    }
    lib.ServerResponse(c, -1, nil)
}

func getUserInfo() error {
    return nil
}
