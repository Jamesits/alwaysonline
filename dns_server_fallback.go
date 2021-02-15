package main

import (
	"github.com/miekg/dns"
	"log"
)

// simply replies NOTIMPL
func handleDefault(this *dnsRequestHandler, r, msg *dns.Msg) {
	log.Printf("[DNS] %d %s not implemented\n", msg.Question[0].Qtype, msg.Question[0].Name)
	msg.Rcode = dns.RcodeNotImplemented
}
