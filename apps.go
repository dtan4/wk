package main

import (
	"net/url"

	"github.com/dtan4/wk/schema"
)

func getApps() (apps []schema.App, err error) {
	token, err := getToken("app.wercker.com")

	if err != nil {
		return nil, err
	}

	client := Client{Endpoint: "https://app.wercker.com/api"}
	opts := url.Values{}
	var appsRes []schema.App
	err = client.Get(&appsRes, "/applications", opts)

	if err != nil {
		return nil, err
	}

	return appsRes, nil
}
