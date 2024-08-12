package dynamicip

import "net"

type IPer interface {
	IP() (net.IP, error)
}
