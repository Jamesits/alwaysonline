package main

import (
	"strconv"
)

func parsePort(service string) (port int, error bool) {
	prt, err := strconv.ParseUint(service, 10, 16)
	if err != nil {
		return 0, true
	}
	return int(prt), false
}
