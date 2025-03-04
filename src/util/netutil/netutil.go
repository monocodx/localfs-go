// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package netutil

import (
	"errors"
	"net"
	"runtime"
	"strings"
)

// implementation based-on chatgpt response
func IPv4Address() ([]string, error) {
	// get network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ipAddrs := []string{}
	// loop interfaces
	for _, iface := range ifaces {
		// filter-out loopback, docker, SIM data, virtualbox and no-up interfaces
		if iface.Name == "lo" ||
			iface.Flags&net.FlagUp == 0 ||
			strings.Contains(iface.Name, "docker") ||
			strings.Contains(iface.Name, "rmnet") ||
			strings.Contains(iface.Name, "dummy") ||
			strings.Contains(iface.Name, "veth") ||
			strings.Contains(iface.Name, "vboxnet") {
			continue
		}

		// get interface associated ip addrs
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		// loop addrs get IPv4 address
		for _, addr := range addrs {
			// cast to *net.IPNet
			ip := addr.(*net.IPNet).IP.To4()
			if ip != nil {
				// filter out Virtualbox Host Only Interface
				// on Windows platform by checking IP 192.168.56.0/24
				// This is not an ideal solution, but it serves as
				// a workaround to address the problem.
				if runtime.GOOS == "windows" &&
					strings.HasPrefix(ip.String(), "192.168.56") {
					continue
				}
				ipAddrs = append(ipAddrs, ip.String())
			}
		}
	}

	if len(ipAddrs) == 0 {
		return nil, errors.New("network interfaces not connected")
	}

	return ipAddrs, nil
}
