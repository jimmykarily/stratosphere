package main

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Db struct {
	connection *sql.DB
	Path       string
}

// Connects to the sqlite database in "path"/
func NewDb(path string) (*Db, error) {
	var err error

	if !fileExists(path) {
		return nil, errors.New("Database file doesn't exist: " + path)
	}

	db := Db{}
	db.connection, err = sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db.Path = path

	return &db, nil
}

// Returns all activities from database as Activity objects
func (db *Db) Activities() ([]*Activity, error) {
	activities := []*Activity{}

	//query := "SELECT * FROM sport_summary ORDER BY start_time DESC;"
	query := "SELECT sport_summary.*, Group_Concat(latitude || ',' || longitude, '|') AS points FROM sport_summary LEFT JOIN location_data ON sport_summary.track_id=location_data.track_id GROUP BY sport_summary.track_id ORDER BY location_data.timestamp ASC;"

	rows, err := db.connection.Query(query)
	if err != nil {
		return activities, err
	}
	defer rows.Close()

	for rows.Next() {
		activity := &Activity{}

		//err = rows.Scan(StructForScan(&activity)...)
		err = rows.Scan(&activity.Id, &activity.Type, &activity.Parent,
			&activity.StartTime, &activity.EndTime, &activity.Calories,
			&activity.CurrentStatus, &activity.ContentJSON, &activity.PointsData)
		if err != nil {
			return activities, err
		}

		activity.ParseContent()

		// Fetch activity points
		query = "SELECT latitude || ',' || longitude FROM location_data WHERE track_id=" + strconv.Itoa(activity.Id) + " ORDER BY timestamp ASC"
		pointRows, err := db.connection.Query(query)
		if err != nil {
			return activities, err
		}
		defer pointRows.Close()

		points := []string{}
		var point string
		for pointRows.Next() {
			err = pointRows.Scan(&point)
			if err != nil {
				panic("Couldn't parse point from database: " + err.Error())
			}
			points = append(points, point)
		}
		err = pointRows.Err()
		if err != nil {
			panic(err.Error())
		}
		activity.PointsData = strings.Join(points, "|")

		activities = append(activities, activity)
	}

	err = rows.Err()
	if err != nil {
		return activities, err
	}

	return activities, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// https://stackoverflow.com/a/29185381
// TODO: Consider using sqlx library
// https://github.com/jmoiron/sqlx
func StructForScan(u interface{}) []interface{} {
	val := reflect.ValueOf(u).Elem()
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v[i] = valueField.Addr().Interface()
	}
	return v
}
