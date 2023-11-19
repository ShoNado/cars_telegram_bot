package handleUsersDB

import (
	"cars_telegram_bot/handleCarDB"
	"database/sql"
	"fmt"
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
	Id                int
	IsAdmin           bool
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
	IsInWork          bool
}

func connectDB() {
	cfg := mysql.Config{
		User:   os.Getenv("root@localhost"),
		Passwd: os.Getenv(handleCarDB.GetPass()),
		Net:    "tcp",
		Addr:   handleCarDB.GetAdr(),
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

func AddNewOrder(profile UserProfile) (int, error) {
	connectDB()
	fmt.Println(profile)
	result, err := db.Exec("INSERT INTO users (isadmin, tgid, namefromuser, namefromtg, username, phonenumber, price, brandcountrymodel, engine, transmission, color, other, ordertime, iscompleted, isadminsaw, isinwork) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		false, profile.TgID, profile.NameFromUser, profile.NameFromTg, profile.UserName, profile.PhoneNumber, profile.Price, profile.BrandCountryModel, profile.Engine, profile.Transmission, profile.Color, profile.Other, profile.OrderTime, profile.IsCompleted, false, false)
	fmt.Println(result, err)
	if err != nil {
		return 0, fmt.Errorf("addOrder: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("getID: %v", err)
	}
	return int(id), nil
}

func ShowOrder(id int) (UserProfile, []uint8, error) {
	connectDB()
	var profile UserProfile
	var time1 []uint8
	result := db.QueryRow("SELECT * FROM users WHERE id=? ", id)

	if err := result.Scan(&profile.Id, &profile.IsAdmin, &profile.TgID, &profile.NameFromUser, &profile.NameFromTg, &profile.UserName, &profile.PhoneNumber,
		&profile.Price, &profile.BrandCountryModel, &profile.Engine, &profile.Transmission, &profile.Color,
		&profile.Other, &time1, &profile.IsCompleted, &profile.IsAdminSaw, &profile.IsInWork); err != nil {
		if err == sql.ErrNoRows {
			return profile, time1, fmt.Errorf("ShowOrder %d: no such order", id)
		}
		return profile, time1, fmt.Errorf("ShowOrder %d: %v", id, err)
	}
	return profile, time1, nil
}

func ShowAllOrders() []UserProfile {
	connectDB()
	var ordersList []UserProfile
	ordersListDB, _ := db.Query("SELECT * FROM users WHERE IsCompleted = ?", true)
	for ordersListDB.Next() {
		var profile UserProfile
		var time1 []uint8
		if err := ordersListDB.Scan(&profile.Id, &profile.IsAdmin, &profile.TgID, &profile.NameFromUser, &profile.NameFromTg, &profile.UserName, &profile.PhoneNumber,
			&profile.Price, &profile.BrandCountryModel, &profile.Engine, &profile.Transmission, &profile.Color,
			&profile.Other, &time1, &profile.IsCompleted, &profile.IsAdminSaw, &profile.IsInWork); err != nil {
			return nil
		}
		if profile.IsAdmin == false {
			ordersList = append(ordersList, profile)
		}

	}
	return ordersList
}

func GetTgID(id int) (int, error) {
	connectDB()
	result := db.QueryRow("SELECT * FROM users WHERE id=?", id)
	var profile UserProfile
	var time1 []uint8
	if err := result.Scan(&profile.Id, &profile.IsAdmin, &profile.TgID, &profile.NameFromUser, &profile.NameFromTg, &profile.UserName, &profile.PhoneNumber,
		&profile.Price, &profile.BrandCountryModel, &profile.Engine, &profile.Transmission, &profile.Color,
		&profile.Other, &time1, &profile.IsCompleted, &profile.IsAdminSaw, &profile.IsInWork); err != nil {
		if err == sql.ErrNoRows {
			return int(profile.TgID), fmt.Errorf("GetTgID %d: no such user", id)
		}
		return int(profile.TgID), fmt.Errorf("TgIDsById %d: %v", id, err)
	}

	return int(profile.TgID), nil
}

func GetClientOrder(TgId int) (UserProfile, error) {
	connectDB()
	result := db.QueryRow("SELECT * FROM users WHERE tgId=?", TgId)
	var profile UserProfile
	var time1 []uint8
	if err := result.Scan(&profile.Id, &profile.IsAdmin, &profile.TgID, &profile.NameFromUser, &profile.NameFromTg, &profile.UserName, &profile.PhoneNumber,
		&profile.Price, &profile.BrandCountryModel, &profile.Engine, &profile.Transmission, &profile.Color,
		&profile.Other, &time1, &profile.IsCompleted, &profile.IsAdminSaw, &profile.IsInWork); err != nil {
		if err == sql.ErrNoRows {
			return profile, fmt.Errorf("GetTgID %d: no such user", profile)
		}
		return profile, fmt.Errorf("TgIDsById %d: %v", profile, err)
	}

	return profile, nil
}

func AdminSeen(id int) error {
	connectDB()
	_, err := db.Exec("UPDATE users SET IsAdminSaw = true WHERE id=?", id)
	return err
}

func AdminGotInWork(id int) error {
	connectDB()
	_, err := db.Exec("UPDATE users SET IsInWork = true WHERE id=?", id)
	return err
}
