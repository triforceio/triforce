package command

import (
  "os"
  "fmt"
  "path"
  "strings"
  "io/ioutil"
  "github.com/mitchellh/cli"
  "github.com/triforce.io/triforce/docker"
)

type Init struct {
  Ui cli.Ui
  Client *docker.Docker
}

func defaultName() string {
  wd,_ := os.Getwd()
  return path.Base(wd)
}

func (cmd *Init) Help() string {
  helpText := `
Usage: triforce init [name]

  Builds a new image from a local Dockerfile and starts a container from it.
`
  return strings.TrimSpace(helpText)
}

func (cmd *Init) Run(args []string) int {

  dockerfile,err := ioutil.ReadFile("./Dockerfile")

  if err != nil && os.IsNotExist(err){
    cmd.Ui.Error("There is no Dockerfile to build from in the current directory")
    return 1
  }

  var name string

  if len(args) < 1 {
    name = defaultName()
  } else {
    name = args[0]
  }

  cmd.Ui.Output(fmt.Sprintf("Creating project named: %s", name))

  cmd.Client.Build(dockerfile, name)

  cmd.Client.Start(name)

  return 0
}

func (cmd *Init) Synopsis() string {
  return "Create and start a new container"
}
