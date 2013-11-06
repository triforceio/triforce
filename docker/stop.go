package docker

import (
  "os"
  "fmt"
  "bytes"
  "net/http"
  "io/ioutil"
)

func (client *Docker) Stop(name string) error {
  cwd,_ := os.Getwd()
  cid,err := ioutil.ReadFile(cwd + "/.cidfile")

  if err != nil {
    return err
  }

  url := fmt.Sprintf("%s/containers/%s/stop", host(client.Addr, client.Port), string(cid))
  fmt.Printf("POST %s\n", url)
  data := bytes.NewReader([]byte("{}"))
  resp,err := http.Post(url, "application/json", data)

  if err != nil {
    fmt.Println("Error calling Docker API to start container", err)
  }

  if resp.StatusCode == 404 {
    err = fmt.Errorf("There is no container named %s to stop", name)
    return err
  } else if resp.StatusCode == 500 {
    body,_ := ioutil.ReadAll(resp.Body)
    err = fmt.Errorf("Error stopping container:\n\t%s", string(body))
    return err
  } else {
    return nil
  }

}
