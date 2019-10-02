package main

import (
	"fmt"

	"github.com/zh/mymod"
	"github.com/zh/mymod/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", mymod.Config())
	fmt.Println(serverlib.Hello())
}
