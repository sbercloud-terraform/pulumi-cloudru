package cluster

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru/cce"
)

func CreateCCE(ctx *pulumi.Context, args CCEArgs) (*cce.Cluster, error) {
	clusterRes, err := cce.NewCluster(ctx, "cceCluster", &cce.ClusterArgs{
		ClusterType:          pulumi.String("VirtualMachine"),
		FlavorId:             pulumi.String("cce.s1.small"),
		Name:                 pulumi.String("my-cce"),
		VpcId:                args.VpcId,
		SubnetId:             args.SubnetId,
		ContainerNetworkType: pulumi.String("vpc-router"),
	})
	if err != nil {
		return nil, err
	}

	dataVolumes := make(cce.NodePoolDataVolumeArray, 0)
	dataVolumes = append(dataVolumes, cce.NodePoolDataVolumeArgs{
		Size:       pulumi.Int(100),
		Volumetype: pulumi.String("SAS"),
	})
	_, err = cce.NewNodePool(ctx, "nodePool", &cce.NodePoolArgs{
		RootVolume: &cce.NodePoolRootVolumeArgs{
			Size:       pulumi.Int(40),
			Volumetype: pulumi.String("SAS"),
		},

		DataVolumes: &dataVolumes,

		Os:               pulumi.String("CentOS 7.6"),
		Password:         pulumi.String("Test@1234"), // pls, don't use this password in production :)
		ClusterId:        clusterRes.ID(),
		Name:             pulumi.String("np1"),
		FlavorId:         pulumi.String("c7n.large.4"),
		InitialNodeCount: pulumi.Int(2),
		SubnetId:         args.SubnetId,
	})
	if err != nil {
		return nil, err
	}

	return clusterRes, nil
}
