package util

import (
  "os"
  "path"
)

func defaultName() string {
  wd,_ := os.Getwd()
  return path.Base(wd)
}

func NameFromArgs(args []string) string {

  var name string

  if len(args) < 1 {
    name = defaultName()
  } else {
    name = args[0]
  }

  return name
}
