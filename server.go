package main

import (
	"encoding/json"
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
	ID 		string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Image     string `json:"image"`
	Value     float64 `json:"value"`
}

type TokenHis struct {
	ID 		string `json:"id"`
	ShortName string `json:"shortName"`
	Value     float64 `json:"value"`
	Time  time.Time `json:"time"`
}

type TokenDto struct{
	Token Token `json:"token"`
	TokenHistory []TokenHis `json:"tokenHistory"`
}

var db *gorm.DB

func main() {
	time.Sleep(time.Second * 4)
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	dbObj, err := connectToServer("dndUser", "cB345678", "1433", "dndDb")
	if err != nil {
		fmt.Println(err)
	}
	db = dbObj
	db.AutoMigrate(&Token{})
	db.AutoMigrate(&TokenHis{})
	wg.Add(2)
	
	go func(db *gorm.DB) {
		startServer(db)
		wg.Done()
	}(db)
	//test(db)
	go func(db *gorm.DB) {
		addTokens(db)
		updatePrice(db,5)
		wg.Done()
	}(db)

	wg.Wait()
}

func startServer(db *gorm.DB) {
	fmt.Println("starting server")
	http.Handle("/", http.FileServer(http.Dir("./resources")))
	http.HandleFunc("/tokens/",getAllTokens)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getAllTokens(w http.ResponseWriter, r *http.Request) {
	TokenDtoList := []TokenDto{} 
	allTokens := []Token{}
	db.Find(&allTokens)
	for i := 0; i < len(allTokens); i++ {
		tokensHis := []TokenHis{}
		db.Where("short_name LIKE ?", allTokens[i].ShortName).Order("time").Find(&tokensHis)
		//TokenHisDtoList := []TokenHis{} 
		/*
		for j := 0; j < len(tokensHis); j++{
			append.
			fmt.Println(tokensHis[j])
		}*/
		TokenDtoList = append(TokenDtoList,TokenDto{ Token: allTokens[i], TokenHistory: tokensHis})
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)


    if err := json.NewEncoder(w).Encode(TokenDtoList); err != nil {
        panic(err)
    }

}



func getRandomString(length int) string {
	charString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	str := ""
	for i := 0; i < length; i++ {
		str += string(charString[rand.Intn(len(charString))])
	}
	return str
}
func getRandomNumber() float64{
	return float64(rand.Intn(10)) / 100
}

func connectToServer(dbUser string, dbPass string, port string, dbName string) (*gorm.DB, error) {
	// github.com/denisenkom/go-mssqldb
	fmt.Println("connection to database")
	dsn := "sqlserver://" + dbUser + ":" + dbPass + "@db:" + port + "?database=" + dbName
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
	allTokens := []Token{}
	db.Find(&allTokens)
	if len(allTokens) == 0 {
		fmt.Println("adding tokens")
		db.Create(&Token{ID:getRandomString(16), Name:"Magecoin", ShortName:"MGC", Value:8, Image:"imagetext"})
		db.Create(&Token{ID:getRandomString(16), Name:"Magecoin Goldpiece", ShortName:"MCG", Value:0.08, Image:"imagetext"})
		db.Create(&Token{ID:getRandomString(16), Name:"Mana", ShortName:"MAN", Value:4, Image:"imagetext"}) 
		db.Create(&Token{ID:getRandomString(16), Name:"Goblincoin", ShortName:"GOB", Value:0.02, Image:"imagetext"})
		db.Create(&Token{ID:getRandomString(16), Name:"DnDeChain", ShortName:"DET", Value:0.5, Image:"imagetext"})

		db.Create(&TokenHis{ID:getRandomString(16), ShortName:"MGC", Value:8, Time: time.Now()}) 
		db.Create(&TokenHis{ID:getRandomString(16), ShortName:"MCG", Value:0.08, Time: time.Now()}) 
		db.Create(&TokenHis{ID:getRandomString(16), ShortName:"MAN", Value:4, Time: time.Now()}) 
		db.Create(&TokenHis{ID:getRandomString(16), ShortName:"GOB", Value:0.02, Time: time.Now()})
		db.Create(&TokenHis{ID:getRandomString(16), ShortName:"DET", Value:0.5, Time: time.Now()})
	} else {
		fmt.Println("tokens already present")
	}
}

func updatePrice(db *gorm.DB, min int){
	
	for i := 1; true; i++{
		fmt.Println("updating tokens")
		allTokens := []Token{}
		db.Find(&allTokens)
		for j := 0; j < len(allTokens); j++{
			newValue := calcChangeInCoinPrice(allTokens[j])
			newToken := allTokens[j]
			newToken.Value = newValue
			newTokenHis := TokenHis{ID:getRandomString(16), ShortName:newToken.ShortName, Value:newValue, Time: time.Now()}
			db.Model(allTokens[j]).Update("Value",newValue)
			db.Create(&newTokenHis)
		}
		time.Sleep(time.Minute * 5)
	}	
}

func calcChangeInCoinPrice(token Token) float64{
	value := token.Value
	num := getRandomNumber()
	if rand.Intn(2) == 1 {
		num += 1
	} else {
		num += 0.90
	}
	//fmt.Println(num)
	if num == 0{
		return value
	}
	return (value * num)
}