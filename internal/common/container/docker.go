package container

import (
	"context"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	. "github.com/logrusorgru/aurora"

	"github.com/kuritka/doge-action/internal/common/guards"
)

type Docker struct {
	ctx context.Context
	cli *client.Client
}

func NewDocker(ctx context.Context) *Docker {
	var err error
	d := new(Docker)
	d.ctx = ctx
	d.cli, err = client.NewClientWithOpts(client.FromEnv)
	guards.Must(err, "init docker client")
	return d
}

func (d *Docker) NetworkExists(name string) bool {
	networks, err := d.cli.NetworkList(d.ctx, types.NetworkListOptions{})
	guards.Must(err, "list networks")
	for _, n := range networks {
		if n.Name == name {
			return true
		}
	}
	return false
}

func (d *Docker) NetworkCreate(name string, subnet *net.IPNet) string {
	ipamConfig := new(network.IPAMConfig)
	ipam := new(network.IPAM)
	ipamConfig.Subnet = subnet.String()
	gw := subnet.IP.Mask(subnet.IP.DefaultMask())
	gw[3]++
	ipamConfig.Gateway = gw.String()
	ipam.Driver = "default"
	ipam.Config = append(ipam.Config, *ipamConfig)
	response, err := d.cli.NetworkCreate(d.ctx, name, types.NetworkCreate{IPAM: ipam})
	guards.Must(err, "create network %s (%s)", BrightCyan(name), BrightCyan(subnet.String()))
	return response.ID
}
