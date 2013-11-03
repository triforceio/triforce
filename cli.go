package main

import (
  "github.com/triforce.io/triforce/subcommand"
)

type TriforceCli struct {
}

func (cli TriforceCli) Start(config Config) {
  build := new(Build)
  build.config = config
  subcommand.Parse(build)
}
