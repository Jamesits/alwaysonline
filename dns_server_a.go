package main

import (
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
)

func handleA(this *dnsRequestHandler, r, msg *dns.Msg) {
	log.Printf("[DNS] A %s\n", msg.Question[0].Name)

	switch strings.ToLower(msg.Question[0].Name) {
	case "dns.msftncsi.com.":
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			A:   net.IPv4(131, 107, 255, 255),
		})
		return

	case "resolver1.opendns.com.":
		// for https://github.com/crazy-max/WindowsSpyBlocker/blob/0e48685cf8c2b3f263f4ada9065188d6c9966cac/app/settings.json#L119
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			A:   net.IPv4(208, 67, 222, 222),
		})
		return

	default: // for everything else, resolve to our own IP address
		if localResolveIp4Enabled {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
				A:   localResolveIp4Address,
			})
		} else {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			})
		}
		return
	}
}
