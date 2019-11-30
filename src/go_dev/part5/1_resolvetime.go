package main

import "fmt"

const(
	SecondsPerMinute = 60

	SecondsPerHour = SecondsPerMinute * 60

	SecondsPerDay  = SecondsPerHour *24
)

func resolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondsPerDay

	hour = seconds / SecondsPerHour

	minute = seconds / SecondsPerMinute

	return
}

func main() {
	fmt.Println(resolveTime(1000))

	_,hour,minute := resolveTime(18000)
	fmt.Println(hour,minute)

	day,_,_:=resolveTime(90000)
	fmt.Println(day)
}
