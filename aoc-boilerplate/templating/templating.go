package templating

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/utils"
)

var language string
var templates map[string]string
var fileExtensions map[string]string

func SetLanguage(input string) {
	language = input
}

func GetLanguage() string {
	return language
}

func CreateTemplateFile(directory string) {
	f, err := os.Create(directory + "main." + GetFileExtension())
	utils.Check(err)
	decoded, err := base64.StdEncoding.DecodeString(GetTemplateContents())
	utils.Check(err)

	f.WriteString(string(decoded))
	f.Close()
}

func GetTemplateContents() (contents string) {
	contents = templates[language]
	return
}

func GetFileExtension() (extension string) {
	extension = fileExtensions[language]
	return
}

func LanguageValid(language string) (valid bool) {
	valid = true
	if _, ok := templates[language]; !ok {
		valid = false
	}
	if _, ok := fileExtensions[language]; ok {
		valid = false
	}

	return
}

func Debug() {
	fmt.Printf("Currently configured templates:\n")
	for k, v := range templates {
		fmt.Printf("\t- %s: %s\n", k, v)
	}

	fmt.Printf("\nCurrently configured file extensions:\n")
	for k, v := range fileExtensions {
		fmt.Printf("\t- %s: .%s\n", k, v)
	}
}

func Languages() {
	fmt.Printf("Currently configured languages:\n")
	for k, _ := range templates {
		fmt.Printf("\t- %s\n", k)
	}

	fmt.Printf("\nTo use a configured language:\n\t- pass it in with the --lang flag\n\t- set the AOC_LANG environment variable\n")
}
