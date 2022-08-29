package main

import (
	"fmt"
	"net"
	"os"
	"github.com/go-ping/ping"
)

var url string = "ws.okex.com"
var IPS []string

func get_ip() {
	ips, err := net.LookupIP(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		ip := ip.String() 
		fmt.Printf("%v. IN A %s\n",url, ip)
		IPS = append(IPS, ip)
	}
}

func main() {
	get_ip()
	for _, i := range IPS {
		pinger, err := ping.NewPinger(i)
		if err != nil {
			panic(err)
		}
		pinger.Count = 5
		err = pinger.Run()
		if err != nil {
			panic(err)
		}
		stats := pinger.Statistics()
		fmt.Printf("\n%v", stats)
	}
}
