package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bgentry/speakeasy"
	"github.com/codegangsta/cli"
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

	var email string

	fmt.Printf("Enter email: ")
	_, err := fmt.Scanln(&email)

	switch {
	case err != nil && err.Error() != "unexpected newline":
		fmt.Println(err.Error())
		os.Exit(1)
	case email == "":
		fmt.Println("Email is required")
		os.Exit(1)
	}

	password, err := readPassword("Enter password: ")

	switch {
	case err != nil && err.Error() != "unexpected newline":
		fmt.Println(err.Error())
		os.Exit(1)
	case email == "":
		fmt.Println("Password is required")
		os.Exit(1)
	}

	nrc, err := hkclient.LoadNetRc()

	if err != nil {
		fmt.Println("Error has occured while loading netrc")
		os.Exit(1)
	}

	err = nrc.SaveCreds(hostname, email, token)

	if err != nil {
		fmt.Println("Error has occured while saving netrc")
		os.Exit(1)
	}

	fmt.Println("Logged in.")
}

type Login struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	OauthScope bool   `json:"oauthScope"`
}

func attemptLogin(username, password string) (hostname, token string, err error) {
	// -X POST
	// -H "Content-Type: application/json"
	// -d { "username": username, "password": password, "oauthScope": "cli" }
	// https://app.wercker.com/api/1.0/oauth/basicauthaccesstoken
}

func readPassword(prompt string) (password string, err error) {
	return speakeasy.Ask("Enter password: ")
}

func doLogout(c *cli.Context) {
}

func doTargets(c *cli.Context) {
}
