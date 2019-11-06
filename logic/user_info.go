package logic

import (
	"bytes"
	"fmt"
	"runtime"

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
			fmt.Printf("查询错误: %s\n", err.Error())
			break
		}
		lib.ServerResponse(c, nil, &res)
		return
	}
	lib.ServerResponse(c, err, nil)
}

func getUserInfo(req *GetUserInfoReq, res *GetUserInfoRes) error {
	sql := fmt.Sprintf(`
        SELECT id AS user_id, student AS name, subject, score
        FROM school_table
        WHERE id = %d;
    `, req.UserId)
	db := global.GLogic.Raw(sql).Scan(res)
	if db.Error != nil {
		fmt.Println(Stack(3))
		return db.Error
	}
	return nil
}

func Stack(skip int) string {
	buf := new(bytes.Buffer)

	callers := make([]uintptr, 32)
	n := runtime.Callers(skip, callers)
	frames := runtime.CallersFrames(callers[:n])
	for {
		if f, ok := frames.Next(); ok {
			fmt.Fprintf(buf, "%s\n\t%s:%d (0x%x)\n", f.Function, f.File, f.Line, f.PC)
		} else {
			break
		}
	}
	return buf.String()

}
