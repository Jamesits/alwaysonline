package main

import (
	"github.com/miekg/dns"
	"log"
)

func handleSOA(this *dnsRequestHandler, r, msg *dns.Msg) {
	log.Printf("[DNS] SOA %s\n", msg.Question[0].Name)

	msg.Answer = append(msg.Answer, &dns.SOA{
		Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
		Ns: "ns1.example.com.",
		Mbox: "dnsmaster.example.com.",
		Serial: 114514,
		Refresh: 60,
		Retry: 10,
		Expire: 3600000,
		Minttl: DNSDefaultTTL,
	})
	return
}
