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
	"dao"
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


//var db *gorm.DB = nil
/*
func getDB() *gorm.DB{
	
	if db == nil {
		for i := 0; i < 5; i++ {
			//fmt.Println(db)
			db, err := connectToServer("dndUser", "cB345678", "1433", "dndDb")
			if err != nil {
				fmt.Println("Connetion faild(" + string(i) + ")",err)
			}
			if db != nil {
				db.AutoMigrate(&Token{})
				db.AutoMigrate(&TokenHis{})
				return db
			}
			time.Sleep(time.Second * 2)
		}
	}
	fmt.Println("test")
	return db
}
*/



func main() {
	/*
	fmt.Println(db)
	time.Sleep(time.Second * 8)
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	
	
	wg.Add(2)
	
	go func() {
		startServer()
		wg.Done()
	}()

	go func() {
		addTokens()
		updatePrice(5)
		wg.Done()
	}()

	wg.Wait()
	*/
}

func startServer() {
	fmt.Println("starting server")
	http.Handle("/", http.FileServer(http.Dir("./resources")))
	http.HandleFunc("/tokens/",getAllTokens)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getAllTokens(w http.ResponseWriter, r *http.Request) {
	TokenDtoList := []TokenDto{} 
	allTokens := []Token{}
	getDB().Find(&allTokens)
	for i := 0; i < len(allTokens); i++ {
		tokensHis := []TokenHis{}
		getDB().Where("short_name LIKE ?", allTokens[i].ShortName).Order("time").Find(&tokensHis)
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



func addTokens(){
	allTokens := []Token{}
	getDB().Find(&allTokens)
	if len(allTokens) == 0 {
		fmt.Println("adding tokens")
		getDB().Create(&Token{ID:getRandomString(16), Name:"Magecoin", ShortName:"MGC", Value:8, Image:"imagetext"})
		getDB().Create(&Token{ID:getRandomString(16), Name:"Magecoin Goldpiece", ShortName:"MCG", Value:0.08, Image:"imagetext"})
		getDB().Create(&Token{ID:getRandomString(16), Name:"Mana", ShortName:"MAN", Value:4, Image:"imagetext"}) 
		getDB().Create(&Token{ID:getRandomString(16), Name:"Goblincoin", ShortName:"GOB", Value:0.02, Image:"imagetext"})
		getDB().Create(&Token{ID:getRandomString(16), Name:"DnDeChain", ShortName:"DET", Value:0.5, Image:"imagetext"})

		getDB().Create(&TokenHis{ID:getRandomString(16), ShortName:"MGC", Value:8, Time: time.Now()}) 
		getDB().Create(&TokenHis{ID:getRandomString(16), ShortName:"MCG", Value:0.08, Time: time.Now()}) 
		getDB().Create(&TokenHis{ID:getRandomString(16), ShortName:"MAN", Value:4, Time: time.Now()}) 
		getDB().Create(&TokenHis{ID:getRandomString(16), ShortName:"GOB", Value:0.02, Time: time.Now()})
		getDB().Create(&TokenHis{ID:getRandomString(16), ShortName:"DET", Value:0.5, Time: time.Now()})
	} else {
		fmt.Println("tokens already present")
	}
}

func updatePrice(min int){
	
	for i := 1; true; i++{
		fmt.Println("updating tokens")
		allTokens := []Token{}
		getDB().Find(&allTokens)
		for j := 0; j < len(allTokens); j++{
			newValue := calcChangeInCoinPrice(allTokens[j])
			newToken := allTokens[j]
			newToken.Value = newValue
			newTokenHis := TokenHis{ID:getRandomString(16), ShortName:newToken.ShortName, Value:newValue, Time: time.Now()}
			getDB().Model(allTokens[j]).Update("Value",newValue)
			getDB().Create(&newTokenHis)
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