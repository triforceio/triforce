package docker

import (
  "os"
  "fmt"
  "bytes"
  "net/http"
  "io/ioutil"
)

func host(addr string, port string) string {
  return fmt.Sprintf("http://%s:%s", addr, port)
}

func (client *Docker) startContainer(cid string, name string) {
  url := fmt.Sprintf("%s/containers/%s/start", host(client.Addr, client.Port), cid)
  fmt.Printf("POST %s\n", url)
  data := bytes.NewReader([]byte("{}"))
  resp,err := http.Post(url, "application/json", data)

  if err != nil {
    fmt.Println("Error calling Docker API to start container", err)
  }

  if resp.StatusCode == 404 {
    fmt.Println("Container does not existing. Creating first.")
    client.Create(name)
    client.startContainer(cid, name)
  } else if resp.StatusCode == 500 {
    body,_ := ioutil.ReadAll(resp.Body)
    fmt.Println("Error starting container", string(body))
  } else {
    fmt.Println("Container started")
  }

}

func (client *Docker) Start(name string) {

  cwd,_ := os.Getwd()
  cid,err := ioutil.ReadFile(cwd + "/.cidfile")

  if os.IsNotExist(err) || string(cid) == "" {
    client.Create(name)
  }

  client.startContainer(string(cid), name)

  if err != nil {
    fmt.Println(err)
    return
  }

}
