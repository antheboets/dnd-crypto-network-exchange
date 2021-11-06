package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
	"net/http"
	"sync"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Token struct {
	ID 		string `db:"id,key,auto"`
	Name      string
	ShortName string
	Image     string
	Value     float64
}


func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	
	wg.Add(1)
	var db *gorm.DB
	fmt.Println(db)
	go func() *gorm.DB {
		db, err := connectToServer("dndUser", "cB345678", "1433", "dndDb")
		if err != nil {
			fmt.Println(err)
		}
		db.AutoMigrate(&Token{})
		wg.Done()
	}()
	fmt.Println(db)
	wg.Wait()
	fmt.Println(db)
	wg.Add(2)
	
	go func() {
		startServer()
		wg.Done()
	}()

	go func() {
		addTokens(nil)
		updatePrice(nil,5)
		wg.Done()
	}()

	wg.Wait()
}

func startServer() {
	fmt.Println("starting server")
	http.Handle("/", http.FileServer(http.Dir("./resources/test")))
	log.Fatal(http.ListenAndServe(":8081", nil))
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

func addTokens(db *gorm.DB){
	fmt.Println("adding tokens")

}

func updatePrice(db *gorm.DB, min int){
	for i := 1; true; i++{
		fmt.Println("updating tokens")
		time.Sleep(time.Minute * 5)
	}
	
}