package main

import (
    "log"
    "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Delete_results(){
  db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
  if err!=nil{
  log.Println("Connection Failed to Open!")
  }
  log.Println("Connection Established Successfully!")
  db.Where("Address=?","Black Pearl").Delete(&Autotable{})
    //We can delete all records using: db.Model(&Autotable{}).Delete(&Autotable{})
  defer db.Close()
}
