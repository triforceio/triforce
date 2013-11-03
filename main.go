package main

/*

  USAGE:

  triforce init [[--vm-path PATH_TO_IMAGE] [--host-name VM_HOSTNAME]]
  triforce build
  triforce ssh

*/


import (
  "fmt"
  "github.com/triforce.io/triforce/subcommand"
)

func main() {
  fmt.Println("main")
  subcommand.Parse(new(subcommand.Build))
}
