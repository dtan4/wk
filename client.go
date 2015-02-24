package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	URL string
}

func (c *Client) Get(v interface{}, path string) error {
	req, err := http.NewRequest("GET", c.URL+path, bytes.NewReader(j))

	if err != nil {
		return err
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&v)

	return err
}

func (c *Client) Post(v interface{}, path string, opts interface{}) error {
	j, err := json.Marshal(opts)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.URL+path, bytes.NewReader(j))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&v)

	return err
}
