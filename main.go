package main

import (
	"github.com/edhuardotierrez/go-myip/myip/cmd"
	"os"
)

func main() {
	if err := cmd.PrintMyIP(); err != nil {
		os.Exit(1)
	}
}
