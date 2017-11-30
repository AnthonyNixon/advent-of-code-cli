package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	"time"
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/base64"
)

var (
	app   = kingpin.New("main", "CLI to boilerplate code for each day of the advent of code")

	newDay = app.Command("get", "Bootstrap a new day for aoc")
	sessionToken = app.Flag("session", "Your session token. Visit https://blog.ajn.me/aoc-session for instructions to get this.").Envar("AOC_SESSION").Required().String()
	lang = app.Flag("lang", "Which language the boilerplate code should be generated in.").Default("go").String()
	year = app.Flag("year", "The year to be used.").Default(fmt.Sprintf("%d", time.Now().Year())).Int()
	day = newDay.Arg("dayNum", "The day to pull inputs for").Default(fmt.Sprintf("%d", time.Now().Day())).Int()
)

const golang  = "cGFja2FnZSBtYWluCgppbXBvcnQgKAogICAgImZtdCIKICAgICJpby9pb3V0aWwiCikKCmZ1bmMgbWFpbigpIHsKICAgIGYsIGVyciA6PSBpb3V0aWwuUmVhZEZpbGUoImlucHV0LnR4dCIpCiAgICBpZiBlcnIgIT0gbmlsIHsKICAgICAgICBmbXQuUHJpbnQoZXJyKQogICAgfQoKICAgIGlucHV0IDo9IHN0cmluZyhmKQp9"
const python = "aW1wb3J0IHJlCgojIFJlYWQgaW4gdGhlIGlucHV0IGZpbGUuCndpdGggb3BlbiAoImlucHV0LnR4dCIsICJyIikgYXMgbXlmaWxlOgogICAgaW5wdXQ9bXlmaWxlLnJlYWQoKQ=="

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case newDay.FullCommand():
		fmt.Printf("Today: %d day %d\n", *year, *day)
		fmt.Printf("Bootstrapping for %s\n", *lang)

		url := fmt.Sprintf("http://adventofcode.com/%d/day/%d/input", *year, *day)
		fmt.Println(url)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			fmt.Println(err.Error())
		}

		cookie := http.Cookie{Name: "session", Value: *sessionToken}
		req.AddCookie(&cookie)
		var client = &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		fmt.Println(string(body))

		dayString := fmt.Sprintf("%d", *day)

		if *day < 10 {
			dayString = "0" + dayString
		}

		directory := fmt.Sprintf("%d/day%s/", *year, dayString)
		fmt.Println("Directory: " + directory)

		dirExists, err := exists(directory)
		if err != nil {
			fmt.Println(err.Error())
		}

		if !dirExists {
			os.MkdirAll(directory, 0777)
		}

		f, err := os.Create(directory + "input.txt")
		if err != nil {
			fmt.Println(err.Error())
		}

		f.WriteString(string(body))
		f.Close()


		switch strings.ToLower(*lang) {
		case "go":
			f, err := os.Create(directory + "main.go")
			if err != nil {
				fmt.Println(err.Error())
			}
			decoded, err := base64.StdEncoding.DecodeString(golang)
			if err != nil {
				fmt.Println(err.Error())
			}

			f.WriteString(string(decoded))
			f.Close()
		case "python":
			f, err := os.Create(directory + "main.py")
			if err != nil {
				fmt.Println(err.Error())
			}
			decoded, err := base64.StdEncoding.DecodeString(python)
			if err != nil {
				fmt.Println(err.Error())
			}

			f.WriteString(string(decoded))
			f.Close()
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}