package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// Badge holds a single badge name & date
type Badge struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

var (
	// LOCATION is the timezone we are in, for use in creating and comparing dates & times
	LOCATION = time.UTC

	// DATE_FORMAT is the format we will use for our dates
	DATE_FORMAT = "2006-01-02"
)

func init() {
}

// Now - returns the time right now when it is called
func Now() time.Time {
	return time.Now().In(LOCATION)
}

// Badges holds all the badges
type Badges struct {
	All []Badge `json:"badges"`
}

// GetBadges is used to unmarshal a JSON badges file into a array of Badge structs
func GetBadges(jsonByteArray []byte) Badges {
	var badges Badges
	json.Unmarshal(jsonByteArray, &badges)
	return badges
}

// checkErr panics if err is not nil
func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	jsonByteArray, err := ioutil.ReadFile("badges.json")
	checkErr(err)
	badges := GetBadges(jsonByteArray)
	numBadges := len(badges.All)
	for i := 0; i < numBadges; i++ {
		d, err := time.Parse(DATE_FORMAT, badges.All[i].Date)
		checkErr(err)
		days := int(time.Since(d).Hours() / 24)
		//weekday := d.Weekday()
		fmt.Printf("%s: %d days", badges.All[i].Name, days)
		if i < numBadges-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(".")
}
