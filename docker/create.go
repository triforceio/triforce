package docker

import (
  "fmt"
  "encoding/json"
  "net"
  "bytes"
  "net/url"
  "net/http"
  "net/http/httputil"
  "io/ioutil"
)

type Container struct {
  Id string
  Hostname string
  Image string
}

func (client *Docker) Create(name string) {
  containerValues := url.Values{}
  containerValues.Set("name", name)

  config := Container{Hostname: name, Image: name}
  body,err := json.Marshal(config)

  req,err := http.NewRequest("POST", "/containers/create?" + containerValues.Encode(), bytes.NewReader(body))

  addr := fmt.Sprintf("%s:%s", client.Addr, client.Port)
  conn,err := net.Dial("tcp", addr)

  if err != nil {
    fmt.Printf("Error dialing Docker API:\n\t\t%s", err)
    return
  }

  clientconn := httputil.NewClientConn(conn, nil)
  resp, err := clientconn.Do(req)
  body,_ = ioutil.ReadAll(resp.Body)

  defer clientconn.Close()

  if err != nil {
    fmt.Printf("Error creating container:\n\t\t%s", err)
    return
  }

  defer resp.Body.Close()

  var c Container

  err = json.Unmarshal(body, &c)

  fmt.Println("Container data: ", c)

  ioutil.WriteFile(".cidfile", []byte(c.Id), 0644)

  fmt.Printf("Created container with id: %s\n", c.Id)
}
