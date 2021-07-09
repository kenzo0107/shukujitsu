package main

import (
	"fmt"
	"time"

	"github.com/kenzo0107/shukujitsu"
)

func main() {
	t, _ := time.Parse("2006-01-02", "2021-07-22")
	if shukujitsu.IsHoliday(t) {
		fmt.Println("there is a holiday")
	}
	fmt.Println(shukujitsu.HolidayName(t))
}
