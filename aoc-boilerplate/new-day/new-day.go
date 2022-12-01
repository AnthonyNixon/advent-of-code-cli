package new_day

import (
	"fmt"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/input"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/templates"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/utils"
)

func NewDay(year int, day int) {
	fmt.Printf("Initializing %d day %d\n", year, day)
	fmt.Printf("Bootstrapping in %s\n", templates.GetLanguage())

	exists, err := utils.CheckIfDirectoryExists(year, day)
	utils.Check(err)

	if exists {
		fmt.Printf("Day already initialized. Please delete the day's directory to re-initialize if necessary.")
		return
	}

	directory := utils.InitializeDirectory(year, day)

	err = input.SaveInputToFile(year, day, directory)
	utils.Check(err)

	templates.CreateTemplateFile(directory)
}
