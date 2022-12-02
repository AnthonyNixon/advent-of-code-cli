package main

import (
	"os"

	new_day "github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/new-day"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/templates"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/token"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/updater"

	"fmt"
	"time"

	"github.com/alecthomas/kingpin"
)

var (
	Version string
	Build   string
	app     = kingpin.New("main", "CLI to boilerplate code for each day of the advent of code")

	version        = app.Command("version", "Version Information")
	template_debug = app.Command("templates", "Shows all current templates configured in the application")

	newDay       = app.Command("get", "Bootstrap a new day for aoc")
	day          = newDay.Arg("dayNum", "The day to pull inputs for").Default(fmt.Sprintf("%d", time.Now().Day())).Int()
	sessionToken = newDay.Flag("session", "Your session token. Visit https://github.com/AnthonyNixon/advent-of-code-boilerplate/blob/main/docs/setup/session.md for instructions.").Envar("AOC_SESSION").Required().String()
	lang         = newDay.Flag("lang", "Which language the boilerplate code should be generated in.").Default("go").String()
	year         = newDay.Flag("year", "The year to be used.").Default(fmt.Sprintf("%d", time.Now().Year())).Int()

	update        = app.Command("update", "Update AOC binary")
	updateVersion = update.Flag("version", "an optional specified version to updater to").Default("latest").String()
)

func init() {
	templates.Initialize()
}

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case version.FullCommand():
		fmt.Printf("%s@%s\n", GetBuild(), GetVersion())
	case newDay.FullCommand():
		token.SetSessionToken(*sessionToken)
		templates.SetLanguage(*lang)

		new_day.NewDay(*year, *day)
	case template_debug.FullCommand():
		templates.Debug()
	case update.FullCommand():
		err := updater.Update(*updateVersion, GetVersion(), GetBuild())
		if err != nil {
			if err.Error() == updater.UpToDateMessage {
				println(updater.UpToDateMessage)
			} else {
				panic(err)
			}
		}
	}
}

func GetVersion() string {
	return Version
}

func GetBuild() string {
	return Build
}
