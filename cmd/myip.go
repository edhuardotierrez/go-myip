package cmd

import (
	"fmt"
	"github.com/edhuardotierrez/go-myip/pkg"
	"github.com/fatih/color"
)

func PrintMyIP() error {

	fmt.Print("\n\n")
	color.Cyan("---------------------------------------------------")
	color.Cyan("Getting your public IP address...")

	n := pkg.New()
	myPublicIp, err := n.GetPublicIP()
	if err != nil {
		fmt.Println("Get Public IP Error:", err.Error())
		return err
	}

	//
	myInterfaceIp, err := n.GetInterfaceIP()
	if err != nil {
		fmt.Println("Get Interface IP Error:", err.Error())
		return err
	}

	hiWhite := color.New(color.FgHiWhite)
	boldHiWhite := hiWhite.Add(color.Bold)

	color.Cyan("Interface: " + myInterfaceIp)
	color.Cyan("---------------------------------------------------")
	color.Cyan("Your Public IP is: \n\n")
	boldHiWhite.Set()
	fmt.Print(myPublicIp, "\n\n")
	color.Unset()
	color.Cyan("---------------------------------------------------")

	return nil
}
