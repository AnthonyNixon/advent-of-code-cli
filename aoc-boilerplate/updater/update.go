package updater

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/inconshreveable/go-update"
)

const apiCall = "https://api.github.com/repos/AnthonyNixon/advent-of-code-boilerplate/releases/latest"
const releaseBase = "https://github.com/AnthonyNixon/advent-of-code-boilerplate/releases/download"
const UpToDateMessage = "Already up to Date."

func getLatestTag() (tag string, err error) {
	type apiResponse struct {
		TagName string `json:"tag_name"`
	}

	resp, err := http.Get(apiCall)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var response apiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	tag = response.TagName
	return
}

func makeUrl(version string, build string) (url string) {
	url = fmt.Sprintf("%s/%s/aoc-%s-%s", releaseBase, version, version, build)
	return
}

func Update(newVer string, currentVer string, build string) (err error) {
	if newVer == "latest" {
		newVer, err = getLatestTag()
		if err != nil {
			return
		}
		fmt.Printf("Latest Tag found: %s\n", newVer)
	}

	fmt.Printf("'%s' == '%s'", newVer, currentVer)
	if newVer == currentVer {
		return errors.New(UpToDateMessage)
	}

	url := makeUrl(newVer, build)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})

	if err != nil {
		return
	}

	fmt.Printf("Successfully updated to version %s\n", newVer)
	return
}
