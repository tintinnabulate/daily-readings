package main

import (
	"fmt"
	"net/http"

	"github.com/tintinnabulate/daily-readings/badges"
	"github.com/tintinnabulate/daily-readings/readings"
)

func badgeHandler(w http.ResponseWriter, r *http.Request) {
	readings.PrintReflections(w)
	fmt.Fprintln(w)
	badges.PrintBadges(w)
}

func main() {
	http.HandleFunc("/daily", badgeHandler)
	http.ListenAndServe(":11000", nil)
}
