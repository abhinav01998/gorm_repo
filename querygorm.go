package main

import (
    "log"
    "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)

func QueryGorm(){
  db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
  if err!=nil{
  log.Println("Connection Failed to Open!")
  }
  log.Println("Connection Established Successfully!")
  user:=&Autotable{Username:"John",Address:"New York"}
  db.First(&Autotable{})
  log.Println(user)
  defer db.Close()
}
