package main

import (
    "log"
    "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Create_insert(){
  db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
  if err!=nil{
  log.Println("Connection Failed to Open!")
  }
  log.Println("Connection Established Successfully!")
    user := &Autotable{Username:"Abhinav", Address:"Los Angeles"}
//  var users []Autotable = []Autotable{
//  UserModel{name: "Ricky",Address:"Sydney"},
//  UserModel{name: "Adam",Address:"Brisbane"},
//  UserModel{name: "Justin",Address:"California"},
//  }
//  for _, user := range users {
//      db.Create(&user)
//  }
    db.Create(user)
    user = &Autotable{Username:"Shahbaaz",Address:"Black Pearl"}
    db.Create(user)
    defer db.Close()
}
