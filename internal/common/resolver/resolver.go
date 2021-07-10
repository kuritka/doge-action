package resolver

import (
	"fmt"
	"net"

	"github.com/AbsaOSS/gopkg/env"
)

type Resolver struct {
}

const (
	defaultNetwork = "k3d-action-bridge-network"
	defaultSubnet  = "172.16.0.0/24"
)

func NewResolver() *Resolver {
	return new(Resolver)
}

func (r *Resolver) Resolve() (o Options, err error) {
	o = Options{Type: K3d}
	var cidrerr error
	o.Network = env.GetEnvAsStringOrFallback(envNetwork, defaultNetwork)
	o.Args = env.GetEnvAsStringOrFallback(envArgs, "")
	o.ClusterName = env.GetEnvAsStringOrFallback(envClusterName, "")
	o.UseRegistry = env.GetEnvAsBoolOrFallback(envDefaultRegistry, false)
	o.RegistryPort, _ = env.GetEnvAsIntOrFallback(envRegistryPort, 5000)
	_, o.Subnet, cidrerr = net.ParseCIDR(env.GetEnvAsStringOrFallback(envSubnetCIDR, defaultSubnet))
	err = r.validate(o, cidrerr)
	o.Verbose = env.GetEnvAsBoolOrFallback(envVerbose, false)
	return
}

func (r *Resolver) validate(o Options, cidrerr error) (err error) {
	if cidrerr != nil {
		return fmt.Errorf("%s: %s", envSubnetCIDR, cidrerr)
	}
	err = field(envNetwork, o.Network).isNotEmpty().err
	if err != nil {
		return err
	}
	o.ClusterName, err = field(envClusterName, o.ClusterName).isNotEmpty().strResult()
	if err != nil {
		return
	}
	if o.UseRegistry {
		err = field(envRegistryPort, o.RegistryPort).isHigherThan(1024).isLessOrEqualTo(65535).err
		if err != nil {
			return
		}
	}
	if o.Network == defaultNetwork && o.Subnet.String() != defaultSubnet {
		return fmt.Errorf("you can't specify custom subnet (%s) for default network (%s)", o.Subnet.String(), o.Network)
	}

	if o.Network != defaultNetwork && o.Subnet.String() == defaultSubnet {
		return fmt.Errorf("subnet CIDR must be specified for custom network (%s)", o.Network)
	}
	return
}
