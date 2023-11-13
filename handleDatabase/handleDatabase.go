package handleDatabase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

type Pass struct {
	Password string
}
type Car struct {
	Id           int
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
	Price        string
	Other        string
	IsCompleted  bool
}

func connectDB() {
	cfg := mysql.Config{
		User:   os.Getenv("root@localhost"),
		Passwd: os.Getenv(getPass()),
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

func ReadAll() ([]Car, error) {
	connectDB()
	var carslist []Car

	carlistDB, err := db.Query("SELECT * FROM cars WHERE IsCompleted = ?", 1)
	if err != nil {
		return nil, fmt.Errorf("access to database: %v", err)
	}
	defer func(carlistDB *sql.Rows) {
		err := carlistDB.Close()
		if err != nil {
			fmt.Printf("allcarslist: %v", err)
		}
	}(carlistDB)

	for carlistDB.Next() {
		var car Car
		if err := carlistDB.Scan(&car.Id, &car.Brand, &car.Model, &car.Country,
			&car.Year, &car.Status, &car.Enginetype, &car.Enginevolume,
			&car.Transmission, &car.DriveType, &car.Color, &car.Mileage,
			&car.Price, &car.Other, &car.IsCompleted); err != nil {
			return nil, fmt.Errorf("getting car list: %v", err)
		}
		carslist = append(carslist, car)
	}
	return carslist, nil
}

func AddNewCar(car Car) (int, error) {
	result, err := db.Exec("INSERT INTO cars (brand, model, country, year, status, enginetype, enginevolume, transmission, drivetype, color, milage, other, IsCompleted) VALUES (?,?,?,?,?,?,?,?,?,?,?, ?, ?)",
		car.Brand, car.Model, car.Country, car.Year, car.Status, car.Enginetype, car.Enginevolume, car.Transmission, car.DriveType, car.Color, car.Mileage, car.Price, car.Other, car.IsCompleted)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return int(id), nil
}
