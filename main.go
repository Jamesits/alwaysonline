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
var disableDNSServer bool
var localDNSPortString *string
var localDNSPort int
var localDNSPortInvaild bool
var localResolveIp4AddressString *string
var localResolveIp4Address net.IP
var localResolveIp4Enabled bool
var localResolveIp6AddressString *string
var localResolveIp6Address net.IP
var localResolveIp6Enabled bool
var localResolvePortString *string
var localResolvePort int
var localResolvePortInvaild bool
var showVersionOnly *bool

func main() {

	// arguments parsing
	flag.BoolVar(&disableDNSServer, "disable-dns-server", false, "disable DNS server")
	localDNSPortString = flag.String("dns-port", "", "listening port of the DNS server")
	localResolveIp4AddressString = flag.String("ipv4", "", "the IPv4 address to this server")
	localResolveIp6AddressString = flag.String("ipv6", "", "the IPv6 address to this server")
	localResolvePortString = flag.String("port", "", "listening port of server")
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

	if len(*localResolvePortString) == 0 {
		localResolvePort = 80
		log.Println("[CONFIG] Listen port: 80")
	} else {
		localResolvePort, localResolvePortInvaild = parsePort(*localResolvePortString)
		if localResolvePortInvaild || localResolvePort == 0 {
			localResolvePort = 80
		}
		log.Printf("[CONFIG] Listen port: %d\n", localResolvePort)
	}

	if len(*localDNSPortString) == 0 {
		localDNSPort = 53
		log.Println("[CONFIG] DNS port: 53")
	} else {
		localDNSPort, localDNSPortInvaild = parsePort(*localDNSPortString)
		if localDNSPortInvaild || localDNSPort == 0 {
			localDNSPort = 53
		}
		log.Printf("[CONFIG] DNS port: %d\n", localDNSPort)
	}

	// HTTP router setup
	mux := http.DefaultServeMux
	mux.HandleFunc("/robots.txt", robots_txt)
	mux.HandleFunc("/ncsi.txt", ncsi_txt)
	mux.HandleFunc("/redirect", redirect)
	mux.HandleFunc("/hotspot-detect.html", hotspot_detect_html)
	mux.HandleFunc("/generate_204", generate_204)
	mux.HandleFunc("/gen_204", generate_204)
	mux.HandleFunc("/nm", nm)
	mux.HandleFunc("/nm-check.txt", nm)
	mux.HandleFunc("/check_network_status.txt", nm)
	mux.HandleFunc("/success.txt", success_txt)
	mux.HandleFunc("/connecttest.txt", connecttest)
	mux.HandleFunc("/connectivity-check.html", connectivity_check_html)
	mux.HandleFunc("/", http_server_fallback) // catch all

	// HTTP logger setup
	loggingHandler := NewApacheLoggingHandler(mux, os.Stdout) // HTTP access log is sent to stdout for now

	// HTTP server setup
	if !(localResolveIp4Enabled || localResolveIp6Enabled) {
		openPlainHttpServer(fmt.Sprintf(":%d", localResolvePort), loggingHandler)
	}
	if localResolveIp4Enabled {
		openPlainHttpServer(fmt.Sprintf("%s:%d", localResolveIp4Address.String(), localResolvePort), loggingHandler)
	}
	if localResolveIp6Enabled {
		openPlainHttpServer(fmt.Sprintf("[%s]:%d", localResolveIp6Address.String(), localResolvePort), loggingHandler)
	}

	// HTTPS server setup
	// tlsHttpServer := &http.Server{
	// 	Addr:     ":443",
	// 	Handler:  loggingHandler,
	// 	TLSConfig: &tls.Config{
	// 		GetCertificate: ,
	// 		// ...
	// 	}
	// }
	// go tlsHttpServer.ListenAndServe()

	if !disableDNSServer {
		// DNS TCP server setup
		dnsTcp1 := &dns.Server{Addr: fmt.Sprintf(":%d", localDNSPort), Net: "tcp"}
		dnsTcp1.Handler = &dnsRequestHandler{}
		go dnsTcp1.ListenAndServe()

		// DNS UDP server setup
		dnsUdp1 := &dns.Server{Addr: fmt.Sprintf(":%d", localDNSPort), Net: "udp"}
		dnsUdp1.Handler = &dnsRequestHandler{}
		go dnsUdp1.ListenAndServe()
	} else {
		log.Println("[MAIN] not starting DNS server")
	}

	// done
	log.Println("[MAIN] Server started.")

	// just a normal while(1)
	mainThreadWaitGroup.Add(1)
	mainThreadWaitGroup.Wait()
}

func openPlainHttpServer(addr string, handler http.Handler) {
	plainHttpServer := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go plainHttpServer.ListenAndServe()
}
