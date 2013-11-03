package main

import (
  "os"
  "os/user"
  "io/ioutil"
  "github.com/BurntSushi/toml"
)

type Config struct {
  DockerApi dockerInfo `toml:"docker"`
}

type dockerInfo struct {
  Host string
  Port string
}

func tryFiles(files ...string) string {
  for _, file := range files {
    contents,err := ioutil.ReadFile(file)
    if !os.IsNotExist(err) {
      return string(contents)
    }
  }
  return ""
}

func ReadConfig() (Config, error) {
  var config Config
  usr,_ := user.Current()
  tomlData := tryFiles(".triforce", usr.HomeDir + "/.triforce")
  if _,err := toml.Decode(tomlData, &config); err != nil {
    return config, err
  }

  return config, nil
}
