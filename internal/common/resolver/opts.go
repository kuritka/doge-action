package resolver

import (
	"net"

	"github.com/AbsaOSS/gopkg/strings"
)

type Provider string

const (
	K3d  Provider = "k3d"
	Kind Provider = "kind"
)

const (
	envClusterName     = "CLUSTER_NAME"
	envArgs            = "ARGS"
	envNetwork         = "NETWORK"
	envSubnetCIDR      = "SUBNET_CIDR"
	envDefaultRegistry = "USE_DEFAULT_REGISTRY"
	envRegistryPort    = "REGISTRY_PORT"
)

// Options defines action input
type Options struct {
	Type Provider
	// Cluster name
	ClusterName string
	// Cluster command line arguments
	Args string
	// Docker bridge network name
	Network string
	// Cluster subnet
	Subnet *net.IPNet
	// Use Registry
	UseRegistry bool
	// Registry port
	RegistryPort int
}

// String prints options as JSON
func (o Options) String() string {
	return strings.ToString(o)
}
