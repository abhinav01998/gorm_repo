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
func Automig() int{
    db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True") //connecting to database
    if err!=nil{
    log.Println("Connection Failed to Open!")
    return 0
    }
    log.Println("Connection Established Successfully!")
    db.Debug().DropTableIfExists(&Autotable{})  //Checking if table altready exists and dropping the table if it does.
    db.Debug().AutoMigrate(&Autotable{})  //Creating the table using automigrate function from the struct Autotable
    defer db.Close()
    return 1
}
