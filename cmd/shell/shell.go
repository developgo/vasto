package shell

import (
	"context"
	"github.com/chrislusf/vasto/client"
)

type ShellOption struct {
	Master     *string
	DataCenter *string
	Keyspace   *string
	Verbose    *bool
}

type shell struct {
	option *ShellOption

	vastoClient *client.VastoClient
}

func RunShell(option *ShellOption) {
	var b = &shell{
		option:      option,
		vastoClient: client.NewClient(context.Background(), "shell", *option.Master, *option.DataCenter),
	}

	b.vastoClient.RegisterForKeyspace(*option.Keyspace)

	if *option.Verbose {
		b.vastoClient.ClusterListener.RegisterShardEventProcessor(b)
		b.vastoClient.ClusterListener.SetVerboseLog(true)
	}

	b.runShell()

}
