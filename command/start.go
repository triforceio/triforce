package command

import (
  "github.com/mitchellh/cli"
  "github.com/triforce.io/triforce/docker"
)

type Start struct {
  Ui cli.Ui
  Client *docker.Docker
}

func (cmd *Start) Help() string {
  return "Usage: triforce start [name]";
}

func (cmd *Start) Synopsis() string {
  return "Start the container for this project"
}

func (cmd *Start) Run(args []string) int {

  var name string

  if len(args) < 1 {
    name = defaultName()
  } else {
    name = args[0]
  }

  cmd.Client.Start(name)

  return 0
}
