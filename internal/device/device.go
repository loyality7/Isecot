package device

import (
	"fmt"
	"net"
	"strings"
)

type Device struct {
	IP        string
	MAC       string
	Hostname  string
	OpenPorts []int
}

func (d *Device) String() string {
	return fmt.Sprintf("IP: %s, MAC: %s, Hostname: %s, Open Ports: %v", d.IP, d.MAC, d.Hostname, d.OpenPorts)
}

func GetDeviceInfo(ip string) (*Device, error) {
	device := &Device{IP: ip}

	// Get MAC address
	ifaces, err := net.Interfaces()
	if err == nil {
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.String() == ip {
					device.MAC = i.HardwareAddr.String()
					break
				}
			}
			if device.MAC != "" {
				break
			}
		}
	}

	// Get hostname
	names, err := net.LookupAddr(ip)
	if err == nil && len(names) > 0 {
		device.Hostname = strings.TrimSuffix(names[0], ".")
	}

	return device, nil
}
