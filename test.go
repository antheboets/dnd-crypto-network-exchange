package main

import(
	"fmt"
	"time"
	"math/rand"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Transaction struct{
	ID string `gorm:"primaryKey"`
	SenderId string `gorm:"size:256"`
	Sender User `gorm:"foreignKey:SenderId"`
	reciverID string `gorm:"size:256"`
	Reciver User `gorm:"foreignKey:reciverID"`
	Amount float64
	TokenID string `gorm:"size:256"`
	Token Token `gorm:"foreignKey:TokenID"`
}

type Token struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	ShortName string  `json:"shortName"`
	Image     string  `json:"image"`
	Value     float64
}

type Wallet struct{
	ID string `json:"id" gorm:"primaryKey"`
	Value float64 `json:"value"`
	Address string 
	TokenID string `gorm:"size:256"`
	Token Token `gorm:"foreignKey:TokenID"`
	UserID string `gorm:"size:256"`
}

type User struct{
	ID string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Wallets []Wallet `json:"wallets" gorm:"foreignKey:UserID"`
}

type TokenHis struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ShortName string    `json:"shortName"`
	Value     float64   `json:"value"`
	Time      time.Time `json:"time"`
	TokenID string `gorm:"size:256"`
	Token Token `gorm:"foreignKey:TokenID"`
}

type TokenDto struct {
	token        Token      `json:"token"`
	TokenHistory []TokenHis `json:"tokenHistory"`
}

type Dao struct{
	db *gorm.DB
}

type Db interface{
	migrateObjects() error
	saveToken(* Token) error
	getToken(ID string,hisLen int) (*TokenDto, error)
}

func (dao *Dao)getToken(ID string,hisLen int) (*TokenDto, error){
	var err error
	var token Token
	var tokenDto TokenDto
	dao.db.Find(&token,"ShortName = ?", ID)
	tokenDto.token = token
	return &tokenDto, err
}

func (dao *Dao)saveToken(token *Token) error{
	var err error
	dao.db.Create(token)
	dao.db.Create(&TokenHis{ID: getRandomString(16),ShortName: token.ShortName, Value: token.Value, Time: time.Now(), TokenID: token.ID})
	return err
}

func (dao *Dao)migrateObjects() error{
	var err error
	dao.db.Migrator().DropTable(&Token{},&TokenHis{},&Wallet{},&User{})
	dao.db.AutoMigrate(&Token{},&TokenHis{},&Wallet{},&User{})
	return err
}

func connectToServer(dbUser string, dbPass string, port string, dbName string) (*gorm.DB, error) {
	// github.com/denisenkom/go-mssqldb
	fmt.Println("connection to database")
	//dsn := "sqlserver://" + dbUser + ":" + dbPass + "@db:" + port + "?database=" + dbName
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

var dao *Dao

func SeedAllTokens() error{
	var err error
	
	tokenList := [...]Token{
		Token{ID: getRandomString(16), Name: "Magecoin", ShortName: "MGC", Value: 8, Image: "imagetext"},
		Token{ID: getRandomString(16), Name: "Magecoin Goldpiece", ShortName: "MCG", Value: 0.08, Image: "imagetext"},
		Token{ID: getRandomString(16), Name: "Mana", ShortName: "MAN", Value: 4, Image: "imagetext"},
		Token{ID: getRandomString(16), Name: "Goblincoin", ShortName: "GOB", Value: 0.02, Image: "imagetext"},
		Token{ID: getRandomString(16), Name: "DnDeChain", ShortName: "DET", Value: 0.5, Image: "imagetext"},
	}
	for _,token := range tokenList{
		dao.saveToken(&token)
	}
	
	return err
}

func getRandomString(length int) string {
	charString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	str := ""
	for i := 0; i < length; i++ {
		str += string(charString[rand.Intn(len(charString))])
	}
	return str
}
func getRandomNumber() float64 {
	return float64(rand.Intn(10)) / 100
}

func main(){
	dbCon, err := connectToServer("dndUser", "cB345678", "1433", "dndDb")
	if err != nil {
		fmt.Println(err)
	}
	dao = &Dao{dbCon}
	dao.migrateObjects()
	fmt.Println("migrate objects")
	SeedAllTokens()
	//fmt.Println(dao.getToken("MGC",10))
}