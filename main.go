package main

/*

  USAGE:

  triforce init [[--vm-path PATH_TO_IMAGE] [--host-name VM_HOSTNAME]]
  triforce build [--docker-host IP_OR_HOSTNAME]
  triforce ssh


  configuration:

  ~/.triforce
  or ./.triforce
  or cli params

  [docker]
  host = 192.168.172.13
  port = 5000 # default is 4243

*/


import (
  "os"
  "fmt"
  "github.com/mitchellh/cli"
)

func main() {
  os.Exit(mainWithReturnCode())
}

func mainWithReturnCode() int {

  // Get the command line args. We shortcut "--version" and "-v" to
  // just show the version.
  args := os.Args[1:]
  for _, arg := range args {
    if arg == "-v" || arg == "--version" {
      newArgs := make([]string, len(args)+1)
      newArgs[0] = "version"
      copy(newArgs[1:], args)
      args = newArgs
      break
    }
  }

  cli := &cli.CLI{
    Args:     args,
    Commands: Commands,
  }

  exitCode, err := cli.Run()

  if err != nil {
          fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
          return 1
  }

  return exitCode

}
