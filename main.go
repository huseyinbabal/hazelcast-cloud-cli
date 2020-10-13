package main

import (
	"github.com/hazelcast/hazelcast-cloud-cli/cmd"
	"github.com/hazelcast/hazelcast-cloud-cli/internal"
	"github.com/hazelcast/hazelcast-cloud-cli/util"
)

func main() {
	cmd.Execute()
	if util.IsCloudShell() {
		return
	}
	updaterService := internal.NewUpdaterService()
	updaterService.Clean()
	updaterService.Check(false)
}
