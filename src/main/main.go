package main

import (
  "serverstuff"
  "os"
  "fmt"
)

func main(){
  names, nameErr := serverstuff.FetchNameList(500)
  if nameErr != nil {
    fmt.Println(nameErr.Error())
    os.Exit(1)
  }

  serverstuff.CreateList(names)
  serverstuff.StartServer()
}
