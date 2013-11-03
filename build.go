package main

import (
  "os"
  "fmt"
  "flag"
  "io/ioutil"
  "github.com/triforce.io/triforce/docker"
)

type Build struct {
  client *docker.Docker
  config Config
}

func (cmd *Build) Name() string {
  return "build";
}

func (cmd *Build) DefineFlags(fs *flag.FlagSet) {
  fmt.Println("got config: ", cmd.config)
  cmd.client = new(docker.Docker)
  cmd.client.Addr = fs.String("docker-api-host", cmd.config.DockerApi.Host, "IP or Hostname of Docker API")
  cmd.client.Port = fs.String("docker-api-port", cmd.config.DockerApi.Port, "Port of Docker API")
}

func (cmd *Build) Run() {

  dockerfile,err := ioutil.ReadFile("./Dockerfile")
  if err != nil {
    if os.IsNotExist(err) {
      err = fmt.Errorf("There is no Dockerfile to build from in the current directory")
    }
    panic(err)
  }

  cmd.client.Build(dockerfile)
}


