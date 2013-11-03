package main

import (
  "fmt"
  "flag"
  "github.com/triforce.io/triforce/subcommand"
)

type build struct {
}

func (cmd *build) Name() string {
  return "build";
}

func (cmd *build) DefineFlags(fs *flag.FlagSet) {
}

func (cmd *build) Run() {
  fmt.Println("build")
}

func main() {
  fmt.Println("main")
  subcommand.Parse(new(build))
}
