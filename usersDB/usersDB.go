package usersDB

import (
	"fmt"

	//"cars_telegram_bot/handleDatabase"
	"database/sql"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

var db *sql.DB

type Pass struct {
	Password string
}

type UserProfile struct {
	id                int
	TgID              int64
	NameFromUser      string
	NameFromTg        string
	UserName          string
	PhoneNumber       string
	Price             string
	BrandCountryModel string
	Engine            string
	Transmission      string
	Color             string
	Other             string
	OrderTime         time.Time
	IsCompleted       bool
	IsAdminSaw        bool
}

func connectDB() {
	cfg := mysql.Config{
		User:   os.Getenv("root@localhost"),
		Passwd: os.Getenv(getPass()),
		Net:    "tcp",
		Addr:   getAdr(),
		DBName: "users",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	db.SetConnMaxLifetime(time.Minute * 3)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
func getPass() string { //read password from file
	file, _ := os.Open("configDB.json")
	decoder := json.NewDecoder(file)
	configuration := Pass{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	return configuration.Password
}

func getAdr() string { //read Ip from file
	file, _ := os.Open("configIP.json")
	decoder := json.NewDecoder(file)
	configuration := Pass{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	return configuration.Password
}

func AddNewOrder(profile UserProfile) (int, error) {
	connectDB()
	result, err := db.Exec("INSERT INTO users (isadmin, tgid, namefromuser, namefromtg, username, phonenumber, price, brandcountrymodel, engine, transmission, color, other, ordertime, iscompleted, isadminsaw, isinwork) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		false, profile.TgID, profile.NameFromUser, profile.NameFromTg, profile.UserName, profile.PhoneNumber, profile.Price, profile.BrandCountryModel, profile.Engine, profile.Transmission, profile.Color, profile.Other, profile.OrderTime, profile.IsCompleted, false, false)
	if err != nil {
		return 0, fmt.Errorf("addOrder: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("getID: %v", err)
	}
	return int(id), nil
}

func ShowOrder(id int) UserProfile {
	var profile UserProfile

	return profile
}
