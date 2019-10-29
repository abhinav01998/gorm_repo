package main

//import "database/sql"
import(
  "log"
  "github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"

)
type Table struct{
    ID int
    username string
}

func GormDemo(){ //15
    db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True") //16
    db.CreateTable(&Table{})
    db.HasTable(&Table{})
    db.DropTableIfExists(&Table{})
    defer db.Close()
    if err != nil {
        log.Println("Connection failed to Open!")
}
        log.Println("Connection Opened Successfully!")
}
