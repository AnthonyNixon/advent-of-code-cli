package main

import (
	"os"

	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/alecthomas/kingpin"
)

var (
	Version string
	app     = kingpin.New("main", "CLI to boilerplate code for each day of the advent of code")

	newDay = app.Command("get", "Bootstrap a new day for aoc")
	day    = newDay.Arg("dayNum", "The day to pull inputs for").Default(fmt.Sprintf("%d", time.Now().Day())).Int()

	sessionToken = app.Flag("session", "Your session token. Visit https://blog.ajn.me/aoc-session for instructions to get this.").Envar("AOC_SESSION").Required().String()
	lang         = app.Flag("lang", "Which language the boilerplate code should be generated in.").Default("go").String()
	year         = app.Flag("year", "The year to be used.").Default(fmt.Sprintf("%d", time.Now().Year())).Int()
	)

func main() {
	app.Version(Version)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case newDay.FullCommand():
		fmt.Printf("Today: %d day %d\n", *year, *day)
		fmt.Printf("Bootstrapping for %s\n", *lang)

		url := fmt.Sprintf("http://adventofcode.com/%d/day/%d/input", *year, *day)
		//fmt.Println(url)
		req, err := http.NewRequest("POST", url, nil)
		check(err)

		cookie := http.Cookie{Name: "session", Value: *sessionToken}
		req.AddCookie(&cookie)
		var client = &http.Client{}
		resp, err := client.Do(req)
		check(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		check(err)

		//fmt.Println(string(body))

		dayString := fmt.Sprintf("%d", *day)

		if *day < 10 {
			dayString = "0" + dayString
		}

		directory := fmt.Sprintf("%d/day%s/", *year, dayString)
		fmt.Println("Directory: " + directory)

		dirExists, err := exists(directory)
		check(err)

		if !dirExists {
			os.MkdirAll(directory, 0777)
			f, err := os.Create(directory + "input.txt")
			check(err)

			f.WriteString(string(body))
			f.Close()

			f, err = os.Create(directory + "input-test.txt")
			check(err)

			f.WriteString("")
			f.Close()

			switch strings.ToLower(*lang) {
			case "go":
				f, err := os.Create(directory + "main.go")
				check(err)
				decoded, err := base64.StdEncoding.DecodeString(golang)
				check(err)

				f.WriteString(string(decoded))
				f.Close()
			case "python":
				f, err := os.Create(directory + "main.py")
				check(err)
				decoded, err := base64.StdEncoding.DecodeString(python)
				check(err)

				f.WriteString(string(decoded))
				f.Close()
			}
		} else {
			fmt.Printf("Day already initialized. Please delete this day's directory to re-initialize if necessary.")
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
