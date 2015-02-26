package main

import (
	"github.com/bgentry/speakeasy"

	"github.com/dtan4/wk/schema"
)

func attemptLogin(username, password string) (hostname, token string, err error) {
	client := Client{Endpoint: "https://app.wercker.com/api"}
	opts := schema.LoginOpts{
		Username:   username,
		Password:   password,
		OauthScope: "cli",
	}
	var loginRes schema.Login
	err = client.Post(&loginRes, "/1.0/oauth/basicauthaccesstoken", opts)

	if err != nil {
		return "", "", err
	}

	// TODO: set hostname dynamically
	return "app.wercker.com", loginRes.Data.AccessToken, nil
}

func getToken(host string) (token string, err error) {
	nrc, err := LoadNetRc()

	if err != nil {
		return "", nil
	}

	m := nrc.FindMachine(host)

	if m == nil {
		return "", nil
	}

	return m.Token, nil
}

func readPassword(prompt string) (password string, err error) {
	return speakeasy.Ask("Enter password: ")
}
