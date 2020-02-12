# TCP-Proxy to fix Docker on windows connectivity issues
#
# From: https://github.com/Microsoft/Docker.DotNet/issues/109

package main

import (
	"log"

	"github.com/google/tcpproxy"
)

func main() {
	var p tcpproxy.Proxy
	p.AddRoute(":2375", tcpproxy.To("0.0.0.0:2375"))
	log.Fatal(p.Run())
}
