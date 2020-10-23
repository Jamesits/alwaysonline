package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"sync"
	"github.com/miekg/dns"
)

var mainThreadWaitGroup = &sync.WaitGroup{}
var localResolveIp4AddressString *string
var localResolveIp4Address net.IP
var localResolveIp4Enabled bool
var localResolveIp6AddressString *string
var localResolveIp6Address net.IP
var localResolveIp6Enabled bool

func main() {
	localResolveIp4AddressString = flag.String("ipv4", "", "the IPv4 address to this server")
	localResolveIp6AddressString = flag.String("ipv6", "", "the IPv6 address to this server")
	flag.Parse()

	if len(*localResolveIp4AddressString) == 0 {
		localResolveIp4Enabled = false
		localResolveIp4Address = net.ParseIP("0.0.0.0")
	} else {
		localResolveIp4Enabled = true
		localResolveIp4Address = net.ParseIP(*localResolveIp4AddressString)
	}

	if len(*localResolveIp6AddressString) == 0 {
		localResolveIp6Enabled = false
		localResolveIp6Address = net.ParseIP("::")
	} else {
		localResolveIp6Enabled = true
		localResolveIp6Address = net.ParseIP(*localResolveIp6AddressString)
	}

	log.Printf("Local server IPv4 address: %s\n", localResolveIp4Address)
	log.Printf("Local server IPv6 address: %s\n", localResolveIp6Address)

	http.HandleFunc("/ncsi.txt", ncsi_txt)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/hotspot-detect.html", hotspot_detect_html)
	http.HandleFunc("/generate_204", generate_204)
	http.HandleFunc("/gen_204", generate_204)
	http.HandleFunc("/nm", nm)
	http.HandleFunc("/success.txt", success_txt)
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

