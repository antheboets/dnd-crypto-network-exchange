package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
  )
  
type User struct{
	Name string
}


type Product struct {
	gorm.Model
	Code  string
	Price uint
  }

func main(){
	  // github.com/denisenkom/go-mssqldb
	  dsn := "sqlserver://dndUser:cB345678@localhost:1433?database=dndDb"
	  db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	  fmt.Println(err)
	  fmt.Println(db)
	  db.AutoMigrate(&Product{})
	  db.Create(&Product{Code: "D42", Price: 100})
	  var product Product
	  db.First(&product, 1) // find product with integer primary key
	  db.First(&product, "code = ?", "D42") // find product with code D42
	
	  // Update - update product's price to 200
	  db.Model(&product).Update("Price", 200)
	  // Update - update multiple fields
	  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	
	  // Delete - delete product
	  //db.Delete(&product, 1)
	  db.First(&product, 2)
	  db.Create(&Product{Code: "D42", Price: 100})
	  fmt.Println(product.Code)
}