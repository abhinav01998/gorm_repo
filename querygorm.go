package main

import (
    "log"
    "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)

func QueryGorm(){
  db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
  db.LogMode(true)
  if err==nil{
  var temp []Autotable
  db.Find(&temp)
  log.Printf("%+v",temp)
  }
  defer db.Close()
}
