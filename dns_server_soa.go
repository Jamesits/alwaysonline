package main

import (
	"github.com/miekg/dns"
	"log"
	"strings"
)

func handleSOA(this *dnsRequestHandler, r, msg *dns.Msg) {
	switch strings.ToLower(msg.Question[0].Name) {
	case ".":
		// When adding an upstream in Windows Server's DNS server, a SOA question to `.` will be generated to probe if the upstream is alive
		// Reply this hardcoded answer to pass the test
		msg.Answer = append(msg.Answer, &dns.SOA{
			Hdr:     dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			Ns:      "a.root-servers.net.",
			Mbox:    "nstld.verisign-grs.com.",
			Serial:  114514,
			Refresh: 60,
			Retry:   10,
			Expire:  3600000,
			Minttl:  DNSDefaultTTL,
		})
		return

	case "dns.msftncsi.com.":
		// in one unknown case, SOA record to dns.msftncsi.com is requested
		msg.Answer = append(msg.Answer, &dns.SOA{
			Hdr:     dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
			Ns:      "ns1-205.azure-dns.com.",
			Mbox:    "azuredns-hostmaster.microsoft.com.",
			Serial:  1,
			Refresh: 60,
			Retry:   10,
			Expire:  3600000,
			Minttl:  DNSDefaultTTL,
		})
		return

	default:
		handleDefault(this, r, msg)
	}
}
