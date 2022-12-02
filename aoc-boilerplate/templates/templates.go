package templates

import (
	"encoding/base64"
	"fmt"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/utils"
	"os"
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
