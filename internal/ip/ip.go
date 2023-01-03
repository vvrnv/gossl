package ip

import (
	"net"

	"github.com/vvrnv/gossl/internal/log"
)

func GetIPV4(host string) ([]string, error) {

	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, log.Error(err)
	}

	var ipList []string
	for _, ip := range ips {
		if net.ParseIP(ip.String()).To4() != nil {
			ipList = append(ipList, ip.String())
		}
	}

	return ipList, nil
}
