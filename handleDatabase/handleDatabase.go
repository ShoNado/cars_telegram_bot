package handleDatabase

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

type Car struct {
	Brand        string
	Model        string
	Country      string
	Year         int
	Status       string
	Enginetype   string
	Enginevolume float64
	Transmission string
	DriveType    string
	Color        string
	Mileage      float64
	//FavoriteNum  int
	Other string
}

func ConnectDB() {
	cfg := mysql.Config{
		User:   os.Getenv("root@localhost"),
		Passwd: os.Getenv("manman9000"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "cars",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	readAll()

}

func readAll() []Car {
	//var crs Car

	num, _ := db.Query("SELECT COUNT(*) FROM cars")
	fmt.Println(num)
	return nil
}
