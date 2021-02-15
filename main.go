package main

import (
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"net/http"
	"sync"
)

var mainThreadWaitGroup = &sync.WaitGroup{}
var localResolveIp4AddressString *string
var localResolveIp4Address net.IP
var localResolveIp4Enabled bool
var localResolveIp6AddressString *string
var localResolveIp6Address net.IP
var localResolveIp6Enabled bool
var showVersionOnly *bool

func main() {
	localResolveIp4AddressString = flag.String("ipv4", "", "the IPv4 address to this server")
	localResolveIp6AddressString = flag.String("ipv6", "", "the IPv6 address to this server")
	showVersionOnly = flag.Bool("version", false, "show version and quit")
	flag.Parse()

	if *showVersionOnly {
		fmt.Println(getVersionFullString())
		return
	} else {
		log.Println(getVersionFullString())
	}

	if len(*localResolveIp4AddressString) == 0 {
		localResolveIp4Enabled = false
		localResolveIp4Address = net.ParseIP("0.0.0.0")
		log.Println("IPv4 resolution disabled")
	} else {
		localResolveIp4Enabled = true
		localResolveIp4Address = net.ParseIP(*localResolveIp4AddressString)
		log.Printf("Local server IPv4 address: %s\n", localResolveIp4Address)
	}

	if len(*localResolveIp6AddressString) == 0 {
		localResolveIp6Enabled = false
		localResolveIp6Address = net.ParseIP("::")
		log.Println("IPv6 resolution disabled")
	} else {
		localResolveIp6Enabled = true
		localResolveIp6Address = net.ParseIP(*localResolveIp6AddressString)
		log.Printf("Local server IPv6 address: %s\n", localResolveIp6Address)
	}

	http.HandleFunc("/ncsi.txt", ncsi_txt)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/hotspot-detect.html", hotspot_detect_html)
	http.HandleFunc("/generate_204", generate_204)
	http.HandleFunc("/gen_204", generate_204)
	http.HandleFunc("/nm", nm)
	http.HandleFunc("/success.txt", success_txt)
	http.HandleFunc("/connecttest.txt", connecttest)
	http.HandleFunc("/", http_server_fallback) // catch all
	go http.ListenAndServe(":80", nil)

	dnsTcp1 := &dns.Server{Addr: ":53", Net: "tcp"}
	dnsTcp1.Handler = &dnsRequestHandler{}
	go dnsTcp1.ListenAndServe()

	dnsUdp1 := &dns.Server{Addr: ":53", Net: "udp"}
	dnsUdp1.Handler = &dnsRequestHandler{}
	go dnsUdp1.ListenAndServe()

	log.Println("Server started.")

	// just a normal while(1)
	mainThreadWaitGroup.Add(1)
	mainThreadWaitGroup.Wait()
}
