package main

import (
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
)

func handleAAAA(this *dnsRequestHandler, r, msg *dns.Msg) {
	log.Printf("A %s\n", msg.Question[0].Name)

	switch strings.ToLower(msg.Question[0].Name) {
	case "dns.msftncsi.com":
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			A: net.ParseIP("fd3e:4f5a:5b81::1"),
		})
		return

	default: // for everything else, resolve to our own IP address
		if localResolveIp6Enabled {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
				A: localResolveIp6Address,
			})
		} else {
			handleDefault(this, r, msg)
		}
		return
	}
}
