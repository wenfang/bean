package ipaddr

import (
	"net"

	"github.com/pkg/errors"
)

// LocalIPv4 获取本机IPv4地址
func LocalIPv4() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", errors.WithStack(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", errors.WithStack(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("no valid ip address")
}

// LocalIPv4Bytes 获得[]byte表示的本机IPv4地址
func LocalIPv4Bytes() ([]byte, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("no valid ip address")
}
