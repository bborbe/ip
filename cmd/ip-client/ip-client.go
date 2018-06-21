package main

import (
	"os"
	"fmt"
	"flag"
	"errors"
	"net/url"
	"time"
	"runtime"
	"net/http"
	"bytes"
	"io"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var client Client
	flag.StringVar(&client.Url, "url", "", "url to ip server")
	flag.Parse()

	if err := client.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "parameter %s\n", err.Error())
		os.Exit(1)
	}
	ip, err := client.Fetch()
	if err != nil {
		fmt.Fprintf(os.Stderr, "get ip failed: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s %s\n", time.Now().Format("2006-01-02T15:04:05"), ip)
	os.Exit(0)
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

func (c *Client) Fetch() (string, error) {
	resp, err := http.Get(c.Url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return "", err
	}
	b := &bytes.Buffer{}
	if _, err := io.Copy(b, resp.Body); err != nil {
		return "", err
	}
	return b.String(), nil
}
