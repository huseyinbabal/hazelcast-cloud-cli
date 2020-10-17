module github.com/hazelcast/hazelcast-cloud-cli

go 1.15

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/fatih/color v1.7.0
	github.com/hazelcast/hazelcast-cloud-sdk-go v1.0.1
	github.com/jedib0t/go-pretty/v6 v6.0.4
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/spf13/cobra v1.0.0
)

replace github.com/hazelcast/hazelcast-cloud-sdk-go v1.0.1 => github.com/yunussandikci/hazelcast-cloud-sdk-go v0.0.0
