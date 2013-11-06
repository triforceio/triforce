package main

import (
  "os"
  "fmt"
  "github.com/mitchellh/cli"
  "github.com/triforce.io/triforce/docker"
  "github.com/triforce.io/triforce/command"
)

var Commands map[string]cli.CommandFactory

func init() {
  ui := &cli.BasicUi{Writer: os.Stdout}

  config,err := ReadConfig()

  if err != nil {
    ui.Error(fmt.Sprintf("Error reading config: %s", err))
  }

  client := new(docker.Docker)
  client.Addr = config.DockerApi.Host
  client.Port = config.DockerApi.Port

  Commands = map[string]cli.CommandFactory {
    "init": func() (cli.Command, error) {
      return &command.Init{
        Client: client,
        Ui:     ui,
      }, nil
    },
    "start": func() (cli.Command, error) {
      return &command.Start{
        Client: client,
        Ui:     ui,
      }, nil
    },
    "stop": func() (cli.Command, error) {
      return &command.Stop{
        Client: client,
        Ui: ui,
      }, nil
    },
  }
}
