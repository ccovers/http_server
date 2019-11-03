package global

import (
    "fmt"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

var GLogic *gorm.DB = nil

func InitLogicDb() error {
    var err error
    for ok := true; ok; ok = false {
        GLogic, err = gorm.Open("mysql",
            "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
        if err != nil {
            break
        }
        GLogic.DB().SetMaxOpenConns(10)
        GLogic.DB().SetMaxIdleConns(3)

        GLogic.LogMode(true)
        //GLogic.SetLogger(clog.Logger)
        fmt.Println("初始化LogicDb成功")
    }
    return err
}

func CloseLogicDb() {
    if GLogic != nil {
        GLogic.Close()
    }
}
