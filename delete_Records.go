package main

import (
    "log"
    "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Delete_results(){
  db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
  if err == nil{
  db.Where("Address=?","Black Pearl").Delete(&Autotable{})
    //We can delete all records using: db.Model(&Autotable{}).Delete(&Autotable{})
  }else{
    log.Println("Cannot delete entry")
  }
  defer db.Close()
}
