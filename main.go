package main

var startFromDay = 1

var adventCalendar []func() error = []func() error{
	day1,
}

func main() {
	for _, fn := range adventCalendar[startFromDay-1:] {
		if err := fn(); err != nil {
			panic(err)
		}
	}
}
