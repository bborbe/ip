package main

import (
	"os"
	"fmt"
	"flag"
	"errors"
	"net/url"
)

func main() {
	var client Client
	flag.StringVar(&client.Url, "url", "", "url to ip server")
	flag.Parse()

	if err := client.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "parameter %s\n", err.Error())
		os.Exit(1)
	}

}

type Client struct {
	Url string
}

func (c *Client) Validate() error {
	if c.Url == "" {
		return errors.New("url missing")
	}
	_, err := url.ParseRequestURI(c.Url)
	if err != nil {
		return errors.New("url invalid")
	}
	return nil
}
