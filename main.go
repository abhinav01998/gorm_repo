package main

 import (
 "log"
 "fmt"
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
    connect := Automig() //automigrating the table to database and returning integer value to connect
    if connect == 0 {
      log.Println("Cannot connect to database")
    }
    defer db.Close()
    if err!=nil{  //handling error
      panic(err)
    }
    log.Println(db.HasTable(&Table{}))  //checking if database has the table
    Create_insert()  //calling function to insert values in databse table autotables
    fmt.Println("The entries in database table are: ")
    QueryGorm()  //printing the records of the database table autotables
    var dell int
    log.Println("If you want delete an entry from table, type 1")
    _,err = fmt.Scanf("%d",&dell)
    if err == nil && dell == 1{
      Delete_results()  //calling function to delete the a result from database table
    }
    log.Println("Records after deleting the last record are: ")
    QueryGorm()  //Printing the remaining records in database
}
