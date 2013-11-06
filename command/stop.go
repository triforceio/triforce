package command

import (
  "fmt"
  "github.com/mitchellh/cli"
  "github.com/triforce.io/triforce/docker"
)

type Stop struct {
  Ui cli.Ui
  Client *docker.Docker
}

func (cmd *Stop) Help() string {
  return "Usage: triforce stop [name]";
}

func (cmd *Stop) Synopsis() string {
  return "Stop the container for this project"
}

func (cmd *Stop) Run(args []string) int {

  var name string

  if len(args) < 1 {
    name = defaultName()
  } else {
    name = args[0]
  }

  err := cmd.Client.Stop(name)

  if err != nil {
    cmd.Ui.Error(fmt.Sprintf("Error stopping container named %s:\n\t%s", name, err))
    return 1
  } else {
    cmd.Ui.Output(fmt.Sprintf("Stopped container %s", name))
  }

  return 0
}

