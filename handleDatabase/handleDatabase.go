package handleDatabase

import (
	"database/sql"
	"encoding/json"
	"errors"
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
type Car struct {
	Id           int
	Brand        string
	Model        string
	Country      string
	Year         int
	Status       string
	StatusBool   bool
	Enginetype   string
	Enginevolume string
	Horsepower   string
	Torque       string
	Transmission string
	DriveType    string
	Color        string
	Mileage      string
	Price        string
	Other        string
	IsCompleted  bool
}

func connectDB() {
	cfg := mysql.Config{
		User:   os.Getenv("root@localhost"),
		Passwd: os.Getenv(getPass()),
		Net:    "tcp",
		Addr:   getAdr(),
		DBName: "cars",
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

func ReadAll() ([]Car, error) {
	connectDB()
	var carslist []Car

	carlistDB, err := db.Query("SELECT * FROM cars WHERE IsCompleted = ?", true)
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
			&car.Year, &car.Status, &car.StatusBool, &car.Enginetype, &car.Enginevolume, &car.Horsepower,
			&car.Torque, &car.Transmission, &car.DriveType, &car.Color, &car.Mileage,
			&car.Price, &car.Other, &car.IsCompleted); err != nil {
			return nil, fmt.Errorf("getting car list: %v", err)
		}
		carslist = append(carslist, car)
	}
	return carslist, nil
}

func AddNewCar(newcar Car) (int, error) {
	connectDB()
	result, err := db.Exec("INSERT INTO cars (brand, model, country, year, status, statusBool,enginetype, enginevolume,horsepower, transmission,torque, drivetype, color, milage, price, other, IsCompleted) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		newcar.Brand, newcar.Model, newcar.Country, newcar.Year, newcar.Status, newcar.StatusBool, newcar.Enginetype, newcar.Enginevolume, newcar.Horsepower, newcar.Torque, newcar.Transmission, newcar.DriveType, newcar.Color, newcar.Mileage, newcar.Price, newcar.Other, newcar.IsCompleted)
	if err != nil {
		return 0, fmt.Errorf("addCar: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("getID: %v", err)
	}
	return int(id), nil
}

func ShowCar(id int) (Car, error) {
	connectDB()
	var car Car
	carWithId := db.QueryRow("SELECT * FROM cars WHERE id = ?", id)
	if err := carWithId.Scan(&car.Id, &car.Brand, &car.Model, &car.Country,
		&car.Year, &car.Status, &car.StatusBool, &car.Enginetype, &car.Enginevolume, &car.Horsepower,
		&car.Torque, &car.Transmission, &car.DriveType, &car.Color, &car.Mileage,
		&car.Price, &car.Other, &car.IsCompleted); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return car, fmt.Errorf("ShowCar by ID %d: no such album", id)
		}
		return car, fmt.Errorf("can't get car with id %d: %v", id, err)
	}
	return car, nil
}

func DeleteCar(id int) error {
	connectDB()
	_, err := db.Exec("DELETE FROM cars WHERE id = ?", id)
	return err
}
