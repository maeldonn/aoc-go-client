package aocgoclient

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
)

// AOCClient represents a client for interacting with the Advent of Code website.
// It includes an HTTP client and a session cookie for authentication.
type AOCClient struct {
	cookie string
	client *http.Client
}

// NewClient initializes a new instance of AOCClient.
// It reads the session cookie from the AOC_COOKIE environment variable
// and configures an HTTP client with a cookie jar.
// Returns an error if the cookie jar cannot be created or if the session cookie is missing.
func NewClient() (*AOCClient, error) {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	session, ok := os.LookupEnv("AOC_COOKIE")
	if !ok {
		return nil, fmt.Errorf("AOC_COOKIE env variable not set")
	}

	aocClient := AOCClient{
		cookie: session,
		client: &http.Client{
			Jar:       cookiejar,
			Transport: &http.Transport{},
		},
	}

	return &aocClient, nil
}

// GetInput fetches the puzzle input for the specified year and day from the Advent of Code website.
// It sends an authenticated HTTP GET request using the session cookie and returns the input as a slice of strings.
// Returns an error if the request fails or the input cannot be read.
func (c *AOCClient) GetInput(year, day int) ([]string, error) {
	url := fmt.Sprintf(
		"https://adventofcode.com/%d/day/%d/input",
		year,
		day,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	c.client.Jar.SetCookies(req.URL, []*http.Cookie{
		{
			Name:  "session",
			Value: c.cookie,
		},
	})

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	sc := bufio.NewScanner(resp.Body)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, nil
}
