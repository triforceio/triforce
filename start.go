package main

import (
  "flag"
  "github.com/triforce.io/triforce/docker"
)

type Start struct {
  client *docker.Docker
  config Config
  name *string
}

func (cmd *Start) Name() string {
  return "start";
}

func (cmd *Start) DefineFlags(fs *flag.FlagSet) {
  cmd.client = new(docker.Docker)
  cmd.client.Addr = fs.String("docker-api-host", cmd.config.DockerApi.Host, "IP or Hostname of Docker API")
  cmd.client.Port = fs.String("docker-api-port", cmd.config.DockerApi.Port, "Port of Docker API")
  cmd.name = fs.String("name", defaultName(), "Name for this project's container (must be unique, basename of cwd by default)")
}

func (cmd *Start) Run() {
  cmd.client.Start(*cmd.name)
}
