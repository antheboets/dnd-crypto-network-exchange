package main

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Transaction struct {
	ID             string `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Token          Token
	Amount         float64
	SenderWallet   Wallet
	RecieverWaller Wallet
}

type Token struct {
	ID        string `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	ShortName string
	Image     string
	Value     float64
}

type Wallet struct {
	ID        string `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Value     int
	//UserS      *User  `gorm:"foreignKey:UserId"`
	//TokenS     *Token `gorm:"foreignKey:TokenId"`
	UserObj    User  `gorm:"foreignKey:LocationsID;references:ID"`
	TokenObj   Token `gorm:"foreignKey:LocationID;references:ID"`
	WalletLink string
}

type User struct {
	gorm.Model
	Name string
	//Wallets   []*Wallet
}

/*
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	db, err := connectToServer("dndUser", "cB345678", "1433", "dndDb")
	fmt.Println(err)
	fmt.Println(db)
	//db.AutoMigrate(&Token{})
	//db.AutoMigrate(&User{})
	db.AutoMigrate(&Wallet{})
	db.Create(&User{Name: "anthe"})
	//db.AutoMigrate(&Transaction{})

	/*
		fmt.Println(err)
		fmt.Println(db)
		fmt.Println("%T\n", db)
		db.AutoMigrate(&Product{})
		db.Create(&Product{Code: "D42", Price: 100})
		var product Product
		db.First(&product, 1)                 // find product with integer primary key
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
	*/

}

func getRandomString(length int) string {
	charString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	str := ""
	for i := 0; i < length; i++ {
		str += string(charString[rand.Intn(len(charString))])
	}
	return str
}

func connectToServer(dbUser string, dbPass string, port string, dbName string) (*gorm.DB, error) {
	// github.com/denisenkom/go-mssqldb
	fmt.Println("connection to database")
	dsn := "sqlserver://" + dbUser + ":" + dbPass + "@localhost:" + port + "?database=" + dbName
	fmt.Println(dsn)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("faild to connect to server")
		fmt.Println(err)
		return db, err
	}
	fmt.Println("database connected")
	return db, err
}
