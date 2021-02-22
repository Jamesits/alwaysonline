package main

import (
	"github.com/miekg/dns"
	"log"
)

// all unknown DNS requests are processed here
func handleDefault(this *dnsRequestHandler, r, msg *dns.Msg) {
	log.Printf("[DNS] %d %s not implemented\n", msg.Question[0].Qtype, msg.Question[0].Name)

	if msg.RecursionDesired {
		// Refused
		msg.RecursionAvailable = false
		msg.Rcode = dns.RcodeRefused
	} else {
		// NotImp
		msg.Rcode = dns.RcodeNotImplemented
	}
}
