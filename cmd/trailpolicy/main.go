package main

import (
  "fmt"
  "github.com/rcaught/trailpolicy"
  "io/ioutil"
  "os"
)

func main() {
  bytes, _ := ioutil.ReadAll(os.Stdin)

  if policyDocument, err := trailpolicy.Convert(bytes); err != nil {
    fmt.Print(err.Error())
  } else {
    fmt.Print(policyDocument)
  }
}
