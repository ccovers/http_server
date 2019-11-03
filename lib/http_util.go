package lib

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
)

func ReadRequestBodyJson(c *gin.Context, obj interface{}) error {
    var body []byte
    body, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        return err
    }

    if err := json.Unmarshal(body, obj); err != nil {
        return err
    }

    return nil
}

func ServerResponse(c *gin.Context, err error, data interface{}) {
    res := struct {
        ErrMsg string      `json:"errMsg"`
        Data   interface{} `json:"data"`
    }{}
    if err != nil {
        res.ErrMsg = err.Error()
    }
    res.Data = data
    c.JSON(http.StatusOK, res)
}
