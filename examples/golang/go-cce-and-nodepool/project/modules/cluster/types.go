package cluster

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type ClusterConf struct {
	ClusterType          string       `json:"cluster_type"`
	FlavorId             string       `json:"flavor_id"`
	ContainerNetworkType string       `json:"container_network_type"`
	Name                 string       `json:"name"`
	SubnetCidr           string       `json:"subnetCidr"`
	PoolSize             int          `json:"poolSize"`
	NodePool             NodePoolConf `json:"node_pool_conf"`
}

type NodePoolConf struct {
	InitialNodeCount int            `json:"initial_node_count"`
	Os               string         `json:"os"`
	FlavorId         string         `json:"flavor_id"`
	RootVolume       RootVolumeConf `json:"root_volume_conf"`
	DataVolume       DataVolumeConf `json:"data_volume_conf"` // this should be a slice, if you want more than 1 volume
}

type RootVolumeConf struct {
	Size       int    `json:"size"`
	VolumeType string `json:"volume_type"`
}

type DataVolumeConf struct {
	Size       int    `json:"size"`
	VolumeType string `json:"volume_type"`
}

type CCEArgs struct {
	VpcId      pulumi.IDOutput
	SubnetId   pulumi.IDOutput
	EipId      pulumi.IDInput
	SubnetCidr pulumi.StringInput
	Name       pulumi.StringInput
	NodeCount  pulumi.IntInput
}
