package main

import (
	"fmt"
	"net/http"
)

// http://www.msftncsi.com/ncsi.txt
func ncsi_txt(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Microsoft NCSI")
}

// http://www.msftconnecttest.com/redirect
func redirect(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Location", "http://go.microsoft.com/fwlink/?LinkID=219472&clcid=0x409")
	w.Header().Add("Server", "Microsoft-IIS/114.514")
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte{})
}

// http://captive.apple.com/hotspot-detect.html
func hotspot_detect_html(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>")
}

// http://connectivitycheck.android.com/generate_204
func generate_204(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}

// http://network-test.debian.org/nm
func nm(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("X-NetworkManager-Status", "online")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "NetworkManager is online\n")
}

// http://detectportal.firefox.com/success.txt
func success_txt(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "success\n")
}
