package main

import (
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"net/http"
	"os"
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
		log.Println("[CONFIG] IPv4 resolution disabled")
	} else {
		localResolveIp4Enabled = true
		localResolveIp4Address = net.ParseIP(*localResolveIp4AddressString)
		log.Printf("[CONFIG] Local server IPv4 address: %s\n", localResolveIp4Address)
	}

	if len(*localResolveIp6AddressString) == 0 {
		localResolveIp6Enabled = false
		localResolveIp6Address = net.ParseIP("::")
		log.Println("[CONFIG] IPv6 resolution disabled")
	} else {
		localResolveIp6Enabled = true
		localResolveIp6Address = net.ParseIP(*localResolveIp6AddressString)
		log.Printf("[CONFIG] Local server IPv6 address: %s\n", localResolveIp6Address)
	}

	mux := http.DefaultServeMux
	loggingHandler := NewApacheLoggingHandler(mux, os.Stdout) // HTTP access log is sent to stdout for now
	server := &http.Server{
		Addr:    ":80",
		Handler: loggingHandler,
	}
	mux.HandleFunc("/ncsi.txt", ncsi_txt)
	mux.HandleFunc("/redirect", redirect)
	mux.HandleFunc("/hotspot-detect.html", hotspot_detect_html)
	mux.HandleFunc("/generate_204", generate_204)
	mux.HandleFunc("/gen_204", generate_204)
	mux.HandleFunc("/nm", nm)
	mux.HandleFunc("/check_network_status.txt", nm)
	mux.HandleFunc("/success.txt", success_txt)
	mux.HandleFunc("/connecttest.txt", connecttest)
	mux.HandleFunc("/connectivity-check.html", connectivity_check_html)
	mux.HandleFunc("/", http_server_fallback) // catch all
	go server.ListenAndServe()

	dnsTcp1 := &dns.Server{Addr: ":53", Net: "tcp"}
	dnsTcp1.Handler = &dnsRequestHandler{}
	go dnsTcp1.ListenAndServe()

	dnsUdp1 := &dns.Server{Addr: ":53", Net: "udp"}
	dnsUdp1.Handler = &dnsRequestHandler{}
	go dnsUdp1.ListenAndServe()

	log.Println("[MAIN] Server started.")

	// just a normal while(1)
	mainThreadWaitGroup.Add(1)
	mainThreadWaitGroup.Wait()
}
