// Auto migrating and creating the table in database.
package main

 import (
 "log"
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
 //"github.com/go-sql-driver/mysql"
)

type Autotable struct{
    ID int //'gorm:"primary_key"'
    Username string
    Address string
}
func Automig() {
    db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
    if err!=nil{
    log.Println("Connection Failed to Open!")
    }
    log.Println("Connection Established Successfully!")
    db.Debug().DropTableIfExists(&Autotable{})
    db.Debug().AutoMigrate(&Autotable{})
    defer db.Close()
}
