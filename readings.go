package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type DailyReflection struct {
	Month     string `json:"month"`
	Day       int    `json:"day"`
	Title     string `json:"title"`
	Quotation string `json:"quotation"`
	Citation  string `json:"citation"`
	Reading   string `json:"reading"`
}

// LOCATION is the timezone we are in, for use in creating and comparing dates & times
var LOCATION = time.UTC

type Month int

const (
	Jan Month = 1 + iota
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
	nMonths = int(Dec)
)

var abbrevMonths = map[string]Month{
	"JANUARY":   Jan,
	"FEBRUARY":  Feb,
	"MARCH":     Mar,
	"APRIL":     Apr,
	"MAY":       May,
	"JUNE":      Jun,
	"JULY":      Jul,
	"AUGUST":    Aug,
	"SEPTEMBER": Sep,
	"OCTOBER":   Oct,
	"NOVEMBER":  Nov,
	"DECEMBER":  Dec,
}

var allMonths = make(map[Month]string)

func init() {
	// Populate map
	for k, v := range abbrevMonths {
		allMonths[v] = k
	}
}

func (m Month) String() string {
	return allMonths[m]
}

// Now - returns the time right now when it is called
func Now() time.Time {
	return time.Now().In(LOCATION)
}

// Readings holds the reflections for each property
type Readings struct {
	DailyReadings []DailyReflection `json:"daily_reflections"`
}

// GetReadings is used to unmarshal a JSON reflections file into a array of DailyReflection struct
func GetReadings(jsonByteArray []byte) Readings {
	var reflections Readings
	json.Unmarshal(jsonByteArray, &reflections)
	return reflections
}

// GetTodaysReading gets... today's reading
func GetTodaysReading(readings Readings) DailyReflection {
	now := Now()
	var reading DailyReflection
	for r := range readings.DailyReadings {
		reading := readings.DailyReadings[r]
		monthString := reading.Month
		if month, ok := abbrevMonths[monthString]; ok {
			if time.Month(month) == now.Month() && reading.Day == now.Day() {
				return reading
			}
		}
	}
	return reading
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	jsonByteArray, err := ioutil.ReadFile("reflections.json")
	check(err)
	reflections := GetReadings(jsonByteArray)
	today := GetTodaysReading(reflections)
	fmt.Println(today.Month, today.Day)
	fmt.Println(today.Quotation)
	fmt.Println(today.Citation)
	fmt.Println()
	fmt.Println(today.Reading)
}
