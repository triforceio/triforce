package subcommand

import (
  "fmt"
  "flag"
)

type Build struct {
}

func (cmd *Build) Name() string {
  return "build";
}

func (cmd *Build) DefineFlags(fs *flag.FlagSet) {
}

func (cmd *Build) Run() {
  fmt.Println("build")
}


