package logic

import (
    "fmt"

    "ccovers/http_server/global"
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
    var err error

    for ok := true; ok; ok = false {
        err = lib.ReadRequestBodyJson(c, &req)
        if err != nil {
            fmt.Printf("解析错误: %s\n", err.Error())
            break
        }

        err = getUserInfo(&req, &res)
        if err != nil {
            fmt.Println("查询错误: %s\n", err.Error())
            break
        }
        lib.ServerResponse(c, nil, &res)
    }
    lib.ServerResponse(c, err, nil)
}

func getUserInfo(req *GetUserInfoReq, res *GetUserInfoRes) error {
    sql := fmt.Sprint(`
        SELECT id AS user_id, student AS name, subject, score 
        FROM school_table
        WHERE id = %d;
    `, req.UserId)
    db := global.GLogic.Raw(sql).Scan(res)
    if db.Error != nil {
        return db.Error
    }
    return nil
}
