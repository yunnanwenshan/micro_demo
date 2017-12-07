package main

import (
 "net"
 "fmt"
)

func main() {
    addrs, err := net.InterfaceAddrs()
    fmt.Printf("----------------%v, $v", addrs, err)
}
