package main

import (
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
)

func handleAAAA(this *dnsRequestHandler, r, msg *dns.Msg) {
	switch strings.ToLower(msg.Question[0].Name) {
	case "dns.msftncsi.com.":
		msg.Answer = append(msg.Answer, &dns.AAAA{
			Hdr:  dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			AAAA: net.ParseIP("fd3e:4f5a:5b81::1"),
		})
		return

	case "resolver1.opendns.com.":
		// for https://github.com/crazy-max/WindowsSpyBlocker/blob/0e48685cf8c2b3f263f4ada9065188d6c9966cac/app/settings.json#L119
		msg.Answer = append(msg.Answer, &dns.AAAA{
			Hdr:  dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			AAAA: net.ParseIP("2620:119:35::35"),
		})
		return

	default:
		if localResolveIp6Enabled {
			// for everything else, resolve to our own IP address
			msg.Answer = append(msg.Answer, &dns.AAAA{
				Hdr:  dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
				AAAA: localResolveIp6Address,
			})
		} else {
			// IPv6 not configured, reply empty answer
			msg.Answer = append(msg.Answer, &dns.AAAA{
				Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			})
		}
		return
	}
}
