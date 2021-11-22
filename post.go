package main

import (
	"fmt"

	"github.com/mitchellh/cli"
)

func postPublishCommandFactory(ui cli.Ui) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return &postPublishCommand{
			UI: ui,
		}, nil
	}
}

type postPublishCommand struct {
	UI cli.Ui
}

func (p postPublishCommand) Help() string {
	return "Publish a post to your site."
}

func (p postPublishCommand) Run(args []string) int {
	if len(args) < 1 {
		p.UI.Error("must specify a markdown file to publish")
		return 1
	}
	path := args[0]

	// TODO(paddy): parse markdown(?) file and publish it
	p.UI.Info(fmt.Sprintf("Successfully published post: %+v", path))
	return 0
}

func (p postPublishCommand) Synopsis() string {
	return "Publish a post to a blog."
}

func postRmCommandFactory(ui cli.Ui) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return postRmCommand{
			UI: ui,
		}, nil
	}
}

type postRmCommand struct {
	UI cli.Ui
}

func (p postRmCommand) Help() string {
	return "Remove a post from the blog."
}

func (p postRmCommand) Run(args []string) int {
	p.UI.Info("Successfully removed post.")
	return 0
}

func (p postRmCommand) Synopsis() string {
	return "Remove a post."
}
