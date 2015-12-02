package main

import (
	"fmt"
	"strconv"

	"database/sql"
)

import _ "github.com/go-sql-driver/mysql"

var _connectUser string
var _connectPass string
var _address string
var _database string

func DoesDatabaseExist(address, username, password, database string) bool {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + address + ":3306)/" + database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return false
	}

	return true
}

func CreateDatabase(address, username, password, database string) error {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + address + ":3306)/"

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Creating database ", connectString, database)
	_, err = db.Exec("CREATE DATABASE " + database)
	if err != nil {
		panic(err)
	}

	return nil
}

func DoesTableExist(address, username, password, database, table string) bool {
	_connectUser = username
	_connectPass = password

	connectString := _connectUser + ":" + _connectPass + "@tcp(" + address + ":3306)/" + database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + database)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("SELECT 1 FROM " + table + " LIMIT 1")
	if err != nil {
		return false
	}

	return true
}

func DoCreateTable(address, username, password, database, table string) error {
	_connectUser = username
	_connectPass = password

	connectString := _connectUser + ":" + _connectPass + "@tcp(" + address + ":3306)/" + database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("USE " + database)
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE " + table + " (Id integer NOT NULL, " +
		"  Name varchar(32), " +
		"  City varchar(32), " +
		"  Country varchar(32), " +
		"  IATA varchar(3) NOT NULL PRIMARY KEY, " +
		"  ICAO varchar(4), " +
		"  Lat varchar(32), " +
		"  Lon varchar(32), " +
		"  Altitude varchar(32), " +
		"  TimezoneOffset varchar(32), " +
		"  DST varchar(32), " +
		"  Timezone varchar(1024)) ")
	if err != nil {
		return err
	}

	return nil
}

func CreateTable(address, username, password, database, table string) error {
	_connectUser = username
	_connectPass = password
	_address = address
	_database = database

	if DoesDatabaseExist(address, username, password, database) == false {
		CreateDatabase(address, username, password, database)
	}

	if DoesTableExist(address, username, password, database, table) == false {
		DoCreateTable(address, username, password, database, table)
	}

	return nil
}

func GetAllRows(table string) AirportCodes {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + _address + ":3306)/" + _database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + _database)
	if err != nil {
		panic(err)
	}

	var rv AirportCodes

	rows, err := db.Query("SELECT * FROM " + table)
	for rows.Next() {
		var Id int
		var Name string
		var City string
		var Country string
		var IATA string
		var ICAO string
		var Lat string
		var Lon string
		var Altitude string
		var TimezoneOffset string
		var DST string
		var Timezone string

		err = rows.Scan(&Id, &Name, &City, &Country, &IATA, &ICAO, &Lat, &Lon, &Altitude, &TimezoneOffset, &DST, &Timezone)

		v := AirportCode{Id: strconv.Itoa(Id),
			Name:           Name,
			City:           City,
			Country:        Country,
			IATA:           IATA,
			ICAO:           ICAO,
			Lat:            Lat,
			Lon:            Lon,
			Altitude:       Altitude,
			TimezoneOffset: TimezoneOffset,
			DST:            DST,
			Timezone:       Timezone,
		}

		rv = append(rv, v)
	}

	return rv
}

func GetRows(table, field, value string) AirportCodes {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + _address + ":3306)/" + _database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + _database)
	if err != nil {
		panic(err)
	}

	var rv AirportCodes

	rows, err := db.Query("SELECT * FROM " + table + " WHERE `" + field + "`=\"" + value + "\"")
	for rows.Next() {
		var Id int
		var Name string
		var City string
		var Country string
		var IATA string
		var ICAO string
		var Lat string
		var Lon string
		var Altitude string
		var TimezoneOffset string
		var DST string
		var Timezone string

		err = rows.Scan(&Id, &Name, &City, &Country, &IATA, &ICAO, &Lat, &Lon, &Altitude, &TimezoneOffset, &DST, &Timezone)

		v := AirportCode{Id: strconv.Itoa(Id),
			Name:           Name,
			City:           City,
			Country:        Country,
			IATA:           IATA,
			ICAO:           ICAO,
			Lat:            Lat,
			Lon:            Lon,
			Altitude:       Altitude,
			TimezoneOffset: TimezoneOffset,
			DST:            DST,
			Timezone:       Timezone,
		}

		rv = append(rv, v)
	}

	return rv
}

func AddRow(table string, fields, values []string) error {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + _address + ":3306)/" + _database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("USE " + _database)
	if err != nil {
		return err
	}

	insert := GetInsertCommand(table, fields, values)
	fmt.Println(insert)
	stmt, err := db.Prepare(insert)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		pkField := 4
		update := GetUpdateCommand(table, pkField, fields, values)

		fmt.Println(update)
		stmt, err := db.Prepare(update)
		if err != nil {
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func GetInsertCommand(table string, fields, values []string) string {
	insert := "INSERT INTO " + table + "("
	for i, v := range fields {
		if i != 0 {
			insert += ", "
		}

		insert += v
	}

	insert += ") VALUES ("
	for i, v := range values {
		if i != 0 {
			insert += ", "
		}

		insert += "\"" + v + "\""
	}

	insert += ")"
	return insert
}

func GetUpdateCommand(table string, pkField int, fields, values []string) string {
	update := "UPDATE " + table + " set "
	for i, v := range fields {
		if i == pkField {
			continue
		}

		if i != 0 {
			update += ", "
		}

		update += v + "=\"" + values[i] + "\""
	}

	update += " WHERE "
	update += fields[pkField] + "=\"" + values[pkField] + "\""

	return update
}
