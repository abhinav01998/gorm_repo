package main

 import (
 "log"
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Autotable struct{
    ID int
    Username string
    Address string
}
func Automig() {
    db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
    if err!=nil{
    log.Println("Connection Failed to Open!")
    }
    defer db.Close()
}
