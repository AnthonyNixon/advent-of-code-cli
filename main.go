package main

import (
	new_day "github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/new-day"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/templates"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/token"
	"os"

	"fmt"
	"time"

	"github.com/alecthomas/kingpin"
)

var (
	Version string
	app     = kingpin.New("main", "CLI to boilerplate code for each day of the advent of code")

	version = app.Command("version", "Version Information")

	newDay       = app.Command("get", "Bootstrap a new day for aoc")
	day          = newDay.Arg("dayNum", "The day to pull inputs for").Default(fmt.Sprintf("%d", time.Now().Day())).Int()
	sessionToken = newDay.Flag("session", "Your session token. Visit https://github.com/AnthonyNixon/advent-of-code-boilerplate/blob/main/docs/setup/session.md for instructions.").Envar("AOC_SESSION").Required().String()
	lang         = newDay.Flag("lang", "Which language the boilerplate code should be generated in.").Default("go").String()
	year         = newDay.Flag("year", "The year to be used.").Default(fmt.Sprintf("%d", time.Now().Year())).Int()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case version.FullCommand():
		fmt.Println(Version)
	case newDay.FullCommand():
		token.SetSessionToken(*sessionToken)
		templates.SetLanguage(*lang)
		templates.Initialize()

		new_day.NewDay(*year, *day)
	}

	println()
}
