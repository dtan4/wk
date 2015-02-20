package main

import (
	"log"
	"os"

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
	Usage: "",
	Description: `
`,
	Action: doCreate,
}

var commandStatus = cli.Command{
	Name:  "status",
	Usage: "",
	Description: `
`,
	Action: doStatus,
}

var commandDeploy = cli.Command{
	Name:  "deploy",
	Usage: "",
	Description: `
`,
	Action: doDeploy,
}

var commandBuilds = cli.Command{
	Name:  "builds",
	Usage: "",
	Description: `
`,
	Action: doBuilds,
}

var commandOpen = cli.Command{
	Name:  "open",
	Usage: "",
	Description: `
`,
	Action: doOpen,
}

var commandQueue = cli.Command{
	Name:  "queue",
	Usage: "",
	Description: `
`,
	Action: doQueue,
}

var commandApps = cli.Command{
	Name:  "apps",
	Usage: "",
	Description: `
`,
	Action: doApps,
}

var commandLink = cli.Command{
	Name:  "link",
	Usage: "",
	Description: `
`,
	Action: doLink,
}

var commandLogin = cli.Command{
	Name:  "login",
	Usage: "",
	Description: `
`,
	Action: doLogin,
}

var commandLogout = cli.Command{
	Name:  "logout",
	Usage: "",
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
}

func doLogout(c *cli.Context) {
}

func doTargets(c *cli.Context) {
}
