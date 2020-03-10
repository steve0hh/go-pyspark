package main

import (
  "log"
  "bytes"
  "os"
  "io"
  "io/ioutil"
  "fmt"
  "os/exec"
  "github.com/markbates/pkger"
  "path/filepath"
)


func main() {
  dname, err := ioutil.TempDir("", "a3w012")
  fname := filepath.Join(dname, "main.py")

  f, err := os.Create(fname)

  if err != nil {
    log.Fatal(err)
  }

  defer os.RemoveAll(dname)

  b, err := pkger.Open("/assets/main.py")
  defer b.Close()

  if err != nil {
    log.Fatal(err)
  }

  io.Copy(f, b)
  f.Sync()

  cmd := exec.Command("spark-submit", fname)

  var out bytes.Buffer
  cmd.Stdout = &out
  err = cmd.Run()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("spart-submit output: \n%v", out.String())
}
