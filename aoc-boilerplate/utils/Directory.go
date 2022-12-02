package utils

import (
	"fmt"
	"os"
)

func GetDirectoryName(year int, day int) (directory string) {
	dayString := fmt.Sprintf("%d", day)

	if day < 10 {
		dayString = "0" + dayString
	}

	directory = fmt.Sprintf("%d/day%s/", year, dayString)
	return
}

func CheckIfDirectoryExists(year int, day int) (exists bool, err error) {
	directory := GetDirectoryName(year, day)
	_, err = os.Stat(directory)
	if err == nil {
		exists = true
	}
	if os.IsNotExist(err) {
		exists = false
	}
	return exists, nil
}

func InitializeDirectory(year int, day int) (directory string) {
	directory = GetDirectoryName(year, day)
	err := os.MkdirAll(directory, 0755)
	Check(err)

	// Create input-test.txt file
	f, err := os.Create(directory + "input-test.txt")
	Check(err)

	f.WriteString("")
	f.Close()

	return
}
