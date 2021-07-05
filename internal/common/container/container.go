package container

import "net"

// Container provider
type Container interface {
	NetworkExists(name string) bool

	NetworkCreate(name string, subnet *net.IPNet) string
}
