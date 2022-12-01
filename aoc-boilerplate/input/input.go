package input

import (
	"fmt"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/config"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/token"
	"github.com/anthonynixon/advent-of-code-boilerplate/aoc-boilerplate/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func SaveInputToFile(year int, day int, directory string) (err error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	utils.Check(err)

	sessionToken := token.GetSessionToken()

	cookie := http.Cookie{Name: "session", Value: sessionToken, Domain: ".adventofcode.com", Path: "/"}
	req.AddCookie(&cookie)
	req.Header.Set("User-Agent", config.Useragent)
	var client = &http.Client{}
	resp, err := client.Do(req)
	utils.Check(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.Check(err)
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)

	f, err := os.Create(directory + "input.txt")
	utils.Check(err)

	_, err = f.WriteString(string(body))
	utils.Check(err)

	err = f.Close()
	utils.Check(err)

	return
}
