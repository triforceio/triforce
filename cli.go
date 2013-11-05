package main

import (
  "github.com/triforce.io/triforce/subcommand"
)

type TriforceCli struct {
}

func (cli TriforceCli) Start(config Config) {
  init := new(Init)
  init.config = config

  start := new(Start)
  start.config = config

  subcommand.Parse(init, start)
}
