package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"net"
)

type MyIP struct {
	Protocol   string
	NameServer string
	IPServer   string
}

func New() *MyIP {
	return &MyIP{
		Protocol:   "udp4",
		NameServer: "ns1.google.com:53",
		IPServer:   "o-o.myaddr.l.google.com",
	}
}

// GetInterfaceIP get the ip of your interface, useful when you want to
// get your ip inside a private network, such as wifi network.
func (mi *MyIP) GetInterfaceIP() (string, error) {
	conn, err := net.Dial(mi.Protocol, mi.NameServer)
	if err != nil {
		return "", err
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

// GetPublicIP get the ip that is public to global.
func (mi *MyIP) GetPublicIP() (string, error) {
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, mi.Protocol, mi.NameServer)
		},
	}
	txt, err := r.LookupTXT(context.Background(), mi.IPServer)
	if err != nil {
		return "", err
	}

	if len(txt) == 0 {
		return "", errors.New("[myip] can't get a ip")
	}

	return txt[0], nil
}

func main() {
	n := New()
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
