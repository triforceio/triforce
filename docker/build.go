package docker

import (
  "os"
  "io"
  "io/ioutil"
  "fmt"
  "net"
  "bytes"
  "net/http"
  "net/http/httputil"
  "archive/tar"
)

type archive io.Reader

func buildPackage(content []byte) (archive, error) {
  buf := new(bytes.Buffer)
  tw := tar.NewWriter(buf)

  hdr := &tar.Header{
    Name: "Dockerfile",
    Size: int64(len(content)),
    Typeflag: tar.TypeReg,
  }
  if err := tw.WriteHeader(hdr); err != nil {
    return nil, err
  }
  if _, err := tw.Write(content); err != nil {
    return nil, err
  }

  if err := tw.Close(); err != nil {
    return nil, err
  }
  return buf, nil
}

func (client *Docker) Build(dockerfile []byte) {
  packaged,err := buildPackage(dockerfile)

  if err != nil {
    fmt.Printf("Error making Dockerfile archive:\n\t\t%s", err)
    return
  }

  req,err := http.NewRequest("POST", "/build", packaged)

  req.Host = *client.Addr
  req.Header.Set("Content-Type", "application/tar")

  addr := fmt.Sprintf("%s:%s", *client.Addr, *client.Port)
  conn,err := net.Dial("tcp", addr)

  if err != nil {
    fmt.Printf("Error dialing Docker API:\n\t\t%s", err)
    return
  }

  clientconn := httputil.NewClientConn(conn, nil)
  resp, err := clientconn.Do(req)
  fmt.Println("Made request")
  defer clientconn.Close()

  if err != nil {
    fmt.Printf("Error performing build:\n\t\t%s", err)
    return
  }

  defer resp.Body.Close()

  fmt.Println("Reading response")
  fmt.Println("Got status: ", resp.StatusCode)
  if resp.StatusCode < 200 || resp.StatusCode >= 400 {

    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println("Read response", body)
    if err != nil {
      fmt.Printf("Error parsing response body:\n\t\t%s", err)
      return
    }

    if len(body) == 0 {
      fmt.Println("Empty body from request. Response returned ", http.StatusText(resp.StatusCode))
    }
    fmt.Printf("Error making request. Response returned: %s with body:\n\t\t%s", http.StatusText(resp.StatusCode), body)
  }

  io.Copy(os.Stdout, resp.Body)
}


