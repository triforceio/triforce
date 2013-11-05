package main

import (
  "os"
  "fmt"
  "flag"
  "path"
  "io/ioutil"
  "github.com/triforce.io/triforce/docker"
)

type Init struct {
  client *docker.Docker
  config Config
  name *string
}

func (cmd *Init) Name() string {
  return "init";
}

func (cmd *Init) DefineFlags(fs *flag.FlagSet) {
  fmt.Println("got config: ", cmd.config)
  cmd.client = new(docker.Docker)
  cmd.client.Addr = fs.String("docker-api-host", cmd.config.DockerApi.Host, "IP or Hostname of Docker API")
  cmd.client.Port = fs.String("docker-api-port", cmd.config.DockerApi.Port, "Port of Docker API")
  cmd.name = fs.String("name", defaultName(), "Name for this project's container (must be unique, basename of cwd by default)")
}

func defaultName() string {
  wd,_ := os.Getwd()
  return path.Base(wd)
}

func (cmd *Init) Run() {

  dockerfile,err := ioutil.ReadFile("./Dockerfile")
  if err != nil {
    if os.IsNotExist(err) {
      err = fmt.Errorf("There is no Dockerfile to build from in the current directory")
    }
    panic(err)
  }

  cmd.client.Build(dockerfile, *cmd.name)

  cmd.client.Start(*cmd.name)
}

