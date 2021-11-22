package main

import (
	"fmt"
	"os"

	"darlinggo.co/version"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("tanglesctl", fmt.Sprintf("%s (%s)", version.Tag, version.Hash))
	c.Args = os.Args[1:]
	ui := &cli.ColoredUi{
		InfoColor:  cli.UiColorCyan,
		ErrorColor: cli.UiColorRed,
		WarnColor:  cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}
	c.Commands = map[string]cli.CommandFactory{
		// publish a post
		"posts publish": postPublishCommandFactory(ui),
		"publish":       postPublishCommandFactory(ui),

		// remove a post
		"posts rm":     postRmCommandFactory(ui),
		"posts remove": postRmCommandFactory(ui),
		"posts delete": postRmCommandFactory(ui),

		// sign in
		"auth login": authLoginCommandFactory(ui),
		"login":      authLoginCommandFactory(ui),

		// sign out
		"auth logout": authLogoutCommandFactory(ui),
		"logout":      authLogoutCommandFactory(ui),

		// get auth info
		"auth whoami": authWhoamiCommandFactory(ui),
		"auth info":   authWhoamiCommandFactory(ui),
		"whoami":      authWhoamiCommandFactory(ui),

		// check for updates
		"version": versionCommandFactory(ui),
	}

	c.HiddenCommands = []string{
		"publish",
		"posts remove",
		"posts delete",
		"login",
		"logout",
		"auth info",
		"whoami",
	}

	c.Autocomplete = true

	status, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(status)
}
