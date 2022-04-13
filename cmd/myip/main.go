package main

import (
	"fmt"
	"github.com/edhuardotierrez/go-myip"
	"github.com/fatih/color"
)

func main() {
	n := myip.New()
	myPublicIp, err := n.GetPublicIP()
	if err != nil {
		fmt.Println("Get Public IP Error:", err.Error())
	}

	//
	myInterfaceIp, err := n.GetInterfaceIP()
	if err != nil {
		fmt.Println("Get Interface IP Error:", err.Error())
	}

	hiWhite := color.New(color.FgHiWhite)
	boldHiWhite := hiWhite.Add(color.Bold)

	fmt.Print("\n\n")
	color.Cyan("---------------------------------------------------")
	color.Cyan("Your Public IP (for interface: " + myInterfaceIp + ") is: \n\n")
	boldHiWhite.Set()
	fmt.Print(myPublicIp, "\n\n")
	color.Unset()
}
