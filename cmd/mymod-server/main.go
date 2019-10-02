package main

import (
	"fmt"

	"github.com/zh/mymod"
	"github.com/zh/mymod/internal/auth"
	"github.com/zh/mymod/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", mymod.Config())
	fmt.Println("Auth:", auth.GetAuth())
	fmt.Println(serverlib.Hello())
}
