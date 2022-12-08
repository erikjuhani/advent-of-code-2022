package main

import (
	"fmt"
	"os"
	"time"
)

var adventCalendar []func() error = []func() error{
	Day1, Day2, Day3, Day4, Day5,
	Day6, Day7,
}

func adventDay() int {
	currentDay := time.Now().Day()
	l := len(adventCalendar)

	if currentDay < 26 && currentDay <= l {
		return currentDay
	}

	return l
}

func main() {
	TODAY := os.Getenv("TODAY")
	d := 0

	if TODAY == "1" {
		d = adventDay() - 1
	}

	for _, fn := range adventCalendar[d:] {
		if err := fn(); err != nil {
			panic(err)
		}
		fmt.Println("---")
	}
}
