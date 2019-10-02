package main

import (
	"fmt"

	"github.com/zh/mymod"
	"github.com/zh/mymod/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", mymod.Config())
	fmt.Println(clientlib.Hello())
}
