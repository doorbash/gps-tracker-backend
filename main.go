package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_NAME = "db.sqlite"
)

func index(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	datetime := r.FormValue("datetime")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")
	altitude := r.FormValue("altitude")

	log.Printf("datetime: %s, latitude: %s, longitude: %s, altitude: %s\n", datetime, latitude, longitude, altitude)

	if datetime == "" || latitude == "" || longitude == "" || altitude == "" {
		log.Println("bad input")
		fmt.Fprintf(w, "ERROR")
		return
	}

	dt, err := time.Parse("20060102150405.000", datetime)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	lng, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	alt, err := strconv.ParseFloat(altitude, 64)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", DB_NAME))
	defer db.Close()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	stmt, err := db.Prepare("INSERT INTO LatLng(datetime, lat, lng, alt) values(?,?,?,?)")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	_, err = stmt.Exec(dt.Unix(), lat, lng, alt)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	fmt.Fprintf(w, "OK")
}

func list(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", DB_NAME))
	defer db.Close()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}
	rows, err := db.Query("SELECT * FROM LatLng order by datetime desc limit 100")
	defer rows.Close()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}

	var datetime int64
	var lat float64
	var lng float64
	var alt float64

	ret := make([]LatLng, 0)
	for rows.Next() {
		err = rows.Scan(&datetime, &lat, &lng, &alt)
		ret = append(ret, LatLng{
			Datetime:  datetime,
			Latitude:  lat,
			Longitude: lng,
			ALtitude:  alt,
		})
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "ERROR")
			return
		}
	}
	data, err := json.Marshal(ret)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ERROR")
		return
	}
	w.Write(data)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/", index)
	http.HandleFunc("/list", list)
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
