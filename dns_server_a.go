package main

import (
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
)

func handleA(this *dnsRequestHandler, r, msg *dns.Msg) {
	log.Printf("A %s\n", msg.Question[0].Name)

	switch strings.ToLower(msg.Question[0].Name) {
	case "dns.msftncsi.com":
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			A: net.IPv4(131,107,255,255),
		})
		return

	default: // for everything else, resolve to our own IP address
		if localResolveIp4Enabled {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
				A: localResolveIp4Address,
			})
		} else {
			handleDefault(this, r, msg)
		}
		return
	}
}
