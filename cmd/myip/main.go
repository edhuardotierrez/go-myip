package main

import (
	"github.com/edhuardotierrez/go-myip/myip/pkg"
	"os"
)

func main() {
	if err := pkg.PrintMyIP(); err != nil {
		os.Exit(1)
	}
}
