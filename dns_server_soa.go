package main

import (
	"github.com/miekg/dns"
)

func handleSOA(_ *dnsRequestHandler, r, msg *dns.Msg) {
	// When adding an upstream in Windows Server's DNS server, a SOA question to `.` (or the specific domain, if adding
	// as a conditional forwarder) will be generated to probe if the upstream is alive.
	// Reply this hardcoded answer to pass the test.
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
}
