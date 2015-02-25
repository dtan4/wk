package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	Endpoint string
}

func (c *Client) Post(v interface{}, path string, opts interface{}) error {
	j, err := json.Marshal(opts)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint+path, bytes.NewReader(j))

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&v)

	return err
}
