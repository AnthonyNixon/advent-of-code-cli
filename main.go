package main

import (
	"os"

	new_day "github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/new-day"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/templating"
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
	languages      = app.Command("languages", "Shows all currently configured languages")

	newDay       = app.Command("get", "Bootstrap a new day for aoc")
	day          = newDay.Arg("dayNum", "The day to pull inputs for").Default(fmt.Sprintf("%d", time.Now().Day())).Int()
	sessionToken = newDay.Flag("session", "Your session token. Visit https://github.com/AnthonyNixon/advent-of-code-boilerplate/blob/main/docs/setup/session.md for instructions.").Envar("AOC_SESSION").Required().String()
	lang         = newDay.Flag("lang", "Which language the boilerplate code should be generated in.").Default("go").Envar("AOC_LANG").String()
	year         = newDay.Flag("year", "The year to be used.").Default(fmt.Sprintf("%d", time.Now().Year())).Int()

	update        = app.Command("update", "Update AOC binary")
	updateVersion = update.Arg("version", "an optional specified version to updater to").Default("latest").String()
)

func init() {
	templating.Initialize()
}

func main() {
	command := kingpin.MustParse(app.Parse(os.Args[1:]))
	// Check for new Minor version. Prompt update if there's a new Major/Minor version

	updatemessage := make(chan string)
	go func() {
		if command == update.FullCommand() {
			updatemessage <- ""
			return
		}

		newAvailable, latest := updater.UpdateAvailable(GetVersion())
		if newAvailable {
			updatemessage <- fmt.Sprintf("\nCurrent Version: %s\nNew minor version (%s) available! Run `aoc update` to automatically update\n", GetVersion(), latest)
			return
		}

		updatemessage <- ""
	}()

	switch command {
	case version.FullCommand():
		fmt.Printf("%s@%s\n", GetBuild(), GetVersion())
	case newDay.FullCommand():
		if !templating.LanguageValid(*lang) {
			fmt.Printf("Invalid language '%s'\n", *lang)
			return
		}
		token.SetSessionToken(*sessionToken)
		templating.SetLanguage(*lang)

		new_day.NewDay(*year, *day)
	case template_debug.FullCommand():
		templating.Debug()
	case languages.FullCommand():
		templating.Languages()
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

	msg := <-updatemessage
	fmt.Println(msg)
}

func GetVersion() string {
	return Version
}

func GetBuild() string {
	return Build
}
