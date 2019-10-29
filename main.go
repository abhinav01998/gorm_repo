package main

 import (
 "log"
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
    Automig()
    Create_insert()
    log.Println(db.HasTable(&Table{}))
defer db.Close()
 if err!=nil{
 log.Println("Connection Failed to Open!")
 }
 log.Println("Connection Established Successfully!")
}
