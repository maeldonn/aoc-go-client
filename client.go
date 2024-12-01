package aocgoclient

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
)

type AOCClient struct {
	client *http.Client
}

func NewClient() (*AOCClient, error) {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	aocClient := AOCClient{
		client: &http.Client{
			Jar:       cookiejar,
			Transport: &http.Transport{},
		},
	}

	return &aocClient, nil
}

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
			Value: os.Getenv("AOC_COOKIE"),
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
