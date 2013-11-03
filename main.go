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
  "fmt"
)

func main() {
  config,err := ReadConfig()

  if err != nil {
    fmt.Println("Error reading config: ", err)
  }

  var cli TriforceCli

  cli.Start(config)
}
