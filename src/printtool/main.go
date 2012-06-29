package main

import "fmt"
import "ipp"

func main() {
  var c ipp.CupsServer
  c.SetServer("http://www.google.com")
  t := c.CreateRequest(0x0001)
  m := t.GetMessage()
  fmt.Println(m)
}
