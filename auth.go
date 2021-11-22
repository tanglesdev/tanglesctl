package main

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
)

func authLoginCommandFactory(ui cli.Ui) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return authLoginCommand{
			ui: ui,
		}, nil
	}
}

type authLoginCommand struct {
	ui cli.Ui
}

func (u authLoginCommand) Help() string {
	return `Set up your tangl.es credentials.

Enter the email address you wish to sign in with when prompted. A sign in code
will be sent to that email from auth@tangl.es. Copy the code, and paste it
into the prompt, and tanglesctl will retrieve credentials for your account and
save them to ~/.config/tangl.es/tanglesctl/credentials. The directory will be
created if it doesn't exist yet.`
}

func (u authLoginCommand) Run(args []string) int {
	email, err := u.ui.Ask("Email:")
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		return 1
	}
	email = strings.TrimSpace(email)
	fmt.Printf("Please check %q for your login code.\n", email)
	code, err := u.ui.Ask("Login code:")
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		return 1
	}
	fmt.Printf("Got code %q\n", code)
	return 0
}

func (u authLoginCommand) Synopsis() string {
	return "Set up your tangl.es credentials."
}

func authLogoutCommandFactory(ui cli.Ui) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return authLogoutCommand{
			ui: ui,
		}, nil
	}
}

type authLogoutCommand struct {
	ui cli.Ui
}

func (u authLogoutCommand) Help() string {
	return `Sign out of your tangl.es account.

Running this will remove the credentials file at
~/.config/tangl.es/tanglesctl/credentials and will revoke the refresh token that
was being used.`
}

func (u authLogoutCommand) Run(args []string) int {
	sure, err := u.ui.Ask("Are you sure you wish to log out? [y/n]:")
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		return 1
	}
	sure = strings.ToLower(strings.TrimSpace(sure))
	switch sure {
	case "y", "yes":
	default:
		fmt.Println("Signing out aborted.")
		return 0
	}
	fmt.Println("Signing out...")
	fmt.Println("You are now signed out.")
	return 0
}

func (u authLogoutCommand) Synopsis() string {
	return "Remove your tangl.es credentials."
}

func authWhoamiCommandFactory(ui cli.Ui) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return authWhoamiCommand{
			ui: ui,
		}, nil
	}
}

type authWhoamiCommand struct {
	ui cli.Ui
}

func (u authWhoamiCommand) Help() string {
	return `Display the email address that was used to sign in.`
}

func (u authWhoamiCommand) Run(args []string) int {
	fmt.Println("who knows?")
	return 0
}

func (u authWhoamiCommand) Synopsis() string {
	return "Display the email address that was used to sign in."
}
