package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bgentry/speakeasy"
	"github.com/codegangsta/cli"

	"github.com/dtan4/wk/schema"
)

var Commands = []cli.Command{
	commandCreate,
	commandStatus,
	commandDeploy,
	commandBuilds,
	commandOpen,
	commandQueue,
	commandApps,
	commandLink,
	commandLogin,
	commandLogout,
	commandTargets,
}

var commandCreate = cli.Command{
	Name:  "create",
	Usage: "Create new app",
	Description: `
`,
	Action: doCreate,
}

var commandStatus = cli.Command{
	Name:  "status",
	Usage: "Show app status",
	Description: `
`,
	Action: doStatus,
}

var commandDeploy = cli.Command{
	Name:  "deploy",
	Usage: "Deploy your app",
	Description: `
`,
	Action: doDeploy,
}

var commandBuilds = cli.Command{
	Name:  "builds",
	Usage: "Show builds",
	Description: `
`,
	Action: doBuilds,
}

var commandOpen = cli.Command{
	Name:  "open",
	Usage: "Open app page in browser",
	Description: `
`,
	Action: doOpen,
}

var commandQueue = cli.Command{
	Name:  "queue",
	Usage: "Show build queue",
	Description: `
`,
	Action: doQueue,
}

var commandApps = cli.Command{
	Name:  "apps",
	Usage: "List your apps",
	Description: `
`,
	Action: doApps,
}

var commandLink = cli.Command{
	Name:  "link",
	Usage: "Link your app to wercker",
	Description: `
`,
	Action: doLink,
}

var commandLogin = cli.Command{
	Name:  "login",
	Usage: "Log in to wercker",
	Description: `
`,
	Action: doLogin,
}

var commandLogout = cli.Command{
	Name:  "logout",
	Usage: "Log out from wercker",
	Description: `
`,
	Action: doLogout,
}

var commandTargets = cli.Command{
	Name:  "targets",
	Usage: "",
	Description: `
`,
	Action: doTargets,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doCreate(c *cli.Context) {
}

func doStatus(c *cli.Context) {
}

func doDeploy(c *cli.Context) {
}

func doBuilds(c *cli.Context) {
}

func doOpen(c *cli.Context) {
}

func doQueue(c *cli.Context) {
}

func doApps(c *cli.Context) {
}

func doLink(c *cli.Context) {
}

func doLogin(c *cli.Context) {
	if len(c.Args()) != 0 {
		cli.ShowCommandHelp(c, "login")
		os.Exit(1)
	}

	var username string

	fmt.Printf("Enter username: ")
	_, err := fmt.Scanln(&username)

	switch {
	case err != nil && err.Error() != "unexpected newline":
		fmt.Println(err.Error())
		os.Exit(1)
	case username == "":
		fmt.Println("Username is required")
		os.Exit(1)
	}

	password, err := readPassword("Enter password: ")

	switch {
	case err != nil && err.Error() != "unexpected newline":
		fmt.Println(err.Error())
		os.Exit(1)
	case password == "":
		fmt.Println("Password is required")
		os.Exit(1)
	}

	hostname, token, err := attemptLogin(username, password)

	if err != nil {
		fmt.Println("Error has occured while login: ")
		fmt.Println(err)
		os.Exit(1)
	}

	nrc, err := LoadNetRc()

	if err != nil {
		fmt.Println("Error has occured while loading netrc")
		os.Exit(1)
	}

	err = nrc.SaveCreds(hostname, username, token)

	if err != nil {
		fmt.Println("Error has occured while saving netrc")
		os.Exit(1)
	}

	fmt.Println("Logged in.")
}

func attemptLogin(username, password string) (hostname, token string, err error) {
	client := Client{Endpoint: "https://app.wercker.com/api/1.0"}
	opts := schema.LoginOpts{
		Username:   username,
		Password:   password,
		OauthScope: "cli",
	}
	var loginRes schema.Login
	err = client.Post(&loginRes, "/oauth/basicauthaccesstoken", opts)

	if err != nil {
		return "", "", err
	}

	// TODO: set hostname dynamically
	return "app.wercker.com", loginRes.Data.AccessToken, nil
}

func readPassword(prompt string) (password string, err error) {
	return speakeasy.Ask("Enter password: ")
}

func doLogout(c *cli.Context) {
	if len(c.Args()) != 0 {
		cli.ShowCommandHelp(c, "logout")
		os.Exit(1)
	}

	nrc, err := LoadNetRc()

	if err != nil {
		fmt.Println("Error has occured while loading netrc")
		os.Exit(1)
	}

	// TODO: set hostname dynamically
	host := "app.wercker.com"
	err = nrc.RemoveCreds(host)

	if err != nil {
		fmt.Println("Error has occured while saving netrc")
		os.Exit(1)
	}

	fmt.Println("Logged out.")
}

func doTargets(c *cli.Context) {
}
