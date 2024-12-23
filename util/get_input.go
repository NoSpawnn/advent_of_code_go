package util

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const (
	EnvVarName       = "AOC_SESSION_KEY"                          // Environment variable that holds the session key
	InputBaseUrl     = "https://adventofcode.com/%d/day/%d/input" // Format with year then day
	InputDir         = "inputs/%d"                                // Input directory, format with year
	InputFileNameFmt = "day_%d.txt"                               // Format with day
	UserAgentStr     = "github.com/NoSpawnn/advent_of_code_go"    // Contact info so I can get shouted at for spamming

	Usage = `
Usage:
	$ go run getinput.go <year> <day> <session key>
	`
)

func getSessionKey() (string, error) {
	sessionKey := os.Getenv(EnvVarName)

	if len(sessionKey) == 0 {
		return "", errors.New("session key environment variable is unset")
	}

	return sessionKey, nil
}

func buildRequest(sessionKey string, year, day int) *http.Request {
	reqUrlStr := fmt.Sprintf(InputBaseUrl, year, day)

	req, err := http.NewRequest(http.MethodGet, reqUrlStr, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", UserAgentStr)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionKey,
	})

	return req
}

func download(sessionKey string, year, day int) {
	dir := fmt.Sprintf(InputDir, year)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(dir, os.ModePerm)
	}

	f := filepath.Join(dir, fmt.Sprintf(InputFileNameFmt, day))
	if _, err := os.Stat(f); err == nil {
		log.Fatal("!! input file already exists !!")
		return
	}

	req := buildRequest(sessionKey, year, day)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("DOWNLOAD ERROR %d - requested input does not exist (%d day %d)\n", http.StatusNotFound, year, day)
		return
	} else if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("DOWNLOAD ERROR %d - is your session key correct?\n", http.StatusBadRequest)
		return
	} else if resp.StatusCode != http.StatusOK {
		fmt.Printf("DOWNLOAD ERROR %d - %s", resp.StatusCode, resp.Status)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(f, body, os.ModePerm)
}

func DownloadInput(yearString, dayString string) {
	sessionKey, err := getSessionKey()
	if err != nil {
		if len(os.Args) != 4 {
			fmt.Println(err, Usage)
			return
		}

		sessionKey = os.Args[3]
	}

	year, err := strconv.Atoi(yearString)
	if err != nil {
		log.Fatal(err)
	}

	day, err := strconv.Atoi(dayString)
	if err != nil {
		log.Fatal(err)
	}

	download(sessionKey, year, day)
}
