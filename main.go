package main

var startFromDay = 1

var adventCalendar []func() error = []func() error{
	Day1,
	Day2,
	Day3,
}

func main() {
	for _, fn := range adventCalendar[startFromDay-1:] {
		if err := fn(); err != nil {
			panic(err)
		}
	}
}
