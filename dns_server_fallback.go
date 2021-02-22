package main

import (
	"github.com/miekg/dns"
	"log"
)

// all unknown DNS requests are processed here
func handleDefault(this *dnsRequestHandler, r, msg *dns.Msg) {
	if msg.RecursionDesired {
		// Refused
		msg.RecursionAvailable = false
		msg.Rcode = dns.RcodeRefused

		log.Printf("[DNS] %d %s refused: recursion requested but not available\n", msg.Question[0].Qtype, msg.Question[0].Name)
	} else {
		// NotImp
		msg.Rcode = dns.RcodeNotImplemented

		log.Printf("[DNS] %d %s refused: not implemented\n", msg.Question[0].Qtype, msg.Question[0].Name)
	}
}
