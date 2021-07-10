package create

import (
	"fmt"
	"strings"
	"time"

	"github.com/AbsaOSS/gopkg/shell"
	"github.com/docker/docker/client"
	"github.com/kuritka/doge-action/internal/common"
	"github.com/kuritka/doge-action/internal/common/container"
	. "github.com/logrusorgru/aurora"

	"github.com/kuritka/doge-action/internal/common/guards"
	"github.com/kuritka/doge-action/internal/common/log"
	"github.com/kuritka/doge-action/internal/common/resolver"
)

type Cluster struct {
	opts      resolver.Options
	cli       *client.Client
	container container.Container
}

var logger = log.Log

func NewCluster(container container.Container, opts resolver.Options) (c *Cluster) {
	var err error
	c = new(Cluster)
	c.container = container
	c.opts = opts
	c.cli, err = client.NewClientWithOpts(client.FromEnv)
	guards.Must(err, "init docker client")
	return
}

func (c *Cluster) Run() (err error) {
	var o string
	defer guards.Must(c.cli.Close(), "closing docker client")
	if !c.container.NetworkExists(c.opts.Network) {
		logger.Info().Msgf("%s %s %s", BrightYellow("create new network"), BrightCyan(c.opts.Network), BrightCyan(c.opts.Subnet.String()))
		id := c.container.NetworkCreate(c.opts.Network, c.opts.Subnet)
		logger.Info().Msgf("🕸🧵 network created %s %s", BrightCyan(c.opts.Network), BrightCyan(id))
	}
	logger.Info().Msgf("%s %s", BrightYellow("☕ 🤏 downloading k3d"), BrightCyan(common.K3dVersion))
	o, err = c.downloadK3d()
	logger.Debug().Msg(BrightWhite(o).String())
	guards.Must(err, "downloading k3d")
	logger.Info().Msgf("🍿🤏😋 %s", BrightYellow("installing k3d, the operation may take several seconds..."))
	o, err = c.runK3d(c.opts.Verbose)
	guards.Must(err, o)
	logger.Info().Msgf("\n%s", o)
	logger.Info().Msgf("⌛  %s", BrightYellow("wait until all agents are ready"))
	c.waitForNodesAreReady()
	logger.Info().Msgf("🚀🚀🚀 %s", BrightCyan("DONE"))
	return
}

func (c *Cluster) String() string {
	return "CREATE CLUSTER"
}

func (c *Cluster) downloadK3d() (o string, err error) {
	a := fmt.Sprintf("curl --silent --fail %s | TAG=%s bash", common.K3dURL, common.K3dVersion)
	return command(a)
}

func (c *Cluster) runK3d(verbose bool) (o string, err error) {
	var v string
	if verbose {
		v = "--verbose"
	}
	a := fmt.Sprintf("k3d cluster create %s --wait %s --network %s %s", c.opts.ClusterName, c.opts.Args, c.opts.Network,v)
	return command(a)
}

func (c *Cluster) getNodeStatusList() (statuses []string, err error) {
	var out string
	a := "kubectl get nodes --no-headers | awk '{ print $2}'"
	out, err = command(a)
	if err != nil {
		return
	}
	statuses = strings.Split(out, "\n")
	return
}

func (c *Cluster) waitForNodesAreReady() {
	for i := 0; i < 20; i++ {
		b := true
		statuses, err := c.getNodeStatusList()
		guards.Must(err, "reading node list")
		for _, s := range statuses {
			if s == "NotReady" || s == "" {
				b = false
				break
			}
		}
		if b {
			return
		}
		time.Sleep(time.Second)
	}
}

func command(a string) (o string, err error) {
	cmd := shell.Command{
		Command: "sh",
		Args:    []string{"-c", a},
	}
	o, err = shell.Execute(cmd)
	return
}
