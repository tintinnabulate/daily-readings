package main

import (
	"fmt"
	"net/http"

	"github.com/tintinnabulate/daily_readings/badges"
	"github.com/tintinnabulate/daily_readings/readings"
)

func badgeHandler(w http.ResponseWriter, r *http.Request) {
	readings.PrintReflections(w)
	fmt.Fprintln(w)
	badges.PrintBadges(w)
}

func main() {
	http.HandleFunc("/badges", badgeHandler)
	http.ListenAndServe(":11000", nil)
}
