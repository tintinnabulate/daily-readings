package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Question struct {
	N int    `json:"number"`
	Q string `json:"question"`
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
type Questions struct {
	All []Question `json:"questions"`
}

// GetQuestions is used to unmarshal a JSON questions file into a array of Question structs
func GetQuestions(jsonByteArray []byte) Questions {
	var questions Questions
	json.Unmarshal(jsonByteArray, &questions)
	return questions
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	jsonByteArray, err := ioutil.ReadFile("step10.json")
	check(err)
	questions := GetQuestions(jsonByteArray)
	for i := 0; i < len(questions.All); i++ {
		fmt.Print(questions.All[i].Q)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(" : ")
		reader.ReadString('\n')
	}
}
