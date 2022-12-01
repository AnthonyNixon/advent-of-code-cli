package templates

import (
	"encoding/base64"
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
