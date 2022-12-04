package utils

import (
	"fmt"
	"time"
)

func GetCurrentDay() (day string) {
	loc, err := time.LoadLocation("America/New_York")
	Check(err)

	now := time.Now().In(loc)
	day = fmt.Sprintf("%d", now.Day())
	return
}

func GetCurrentYear() (year string) {

}
