package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	datetime := r.FormValue("datetime")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")
	altitude := r.FormValue("altitude")

	fmt.Fprintf(w, "OK")

	log.Printf("datetime: %s, latitude: %s, longitude: %s, altitude: %s\n", datetime, latitude, longitude, altitude)
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
