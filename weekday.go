package main

import "time"
import "fmt"

func main() {
	layout := "2006-01-02T15:04:05.000Z"
	formStr := "2000-01-01T00:00:00.000Z"
	sobrietyDate, _ := time.Parse(layout, formStr)
	fmt.Printf("%d days sober\n", int(time.Since(sobrietyDate).Hours())/24)
	fmt.Printf("It was a %s\n", sobrietyDate.Weekday())
}
