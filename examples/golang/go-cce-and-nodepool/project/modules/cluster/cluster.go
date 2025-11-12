package cluster

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru/cce"
)

func CreateCCE(ctx *pulumi.Context, args CCEArgs, conf *ClusterConf) (*cce.Cluster, error) {
	// This is used as the internal name of the resource and as the name that is displayed in the console
	clusterName := conf.Name + "-cluster"

	clusterRes, err := cce.NewCluster(ctx, clusterName, &cce.ClusterArgs{
		ClusterType:          pulumi.String(conf.ClusterType),
		FlavorId:             pulumi.String(conf.FlavorId),
		Name:                 pulumi.String(clusterName),
		VpcId:                args.VpcId,
		SubnetId:             args.SubnetId,
		ContainerNetworkType: pulumi.String(conf.ContainerNetworkType),
	})
	if err != nil {
		return nil, err
	}

	// This is used as the internal name of the resource and as the name that is displayed in the console
	nodepoolName := conf.Name + "-nodepool"

	_, err = cce.NewNodePool(ctx, nodepoolName, &cce.NodePoolArgs{
		RootVolume: &cce.NodePoolRootVolumeArgs{
			Size:       pulumi.Int(conf.NodePool.RootVolume.Size),
			Volumetype: pulumi.String(conf.NodePool.RootVolume.VolumeType),
		},
		DataVolumes: cce.NodePoolDataVolumeArray{
			cce.NodePoolDataVolumeArgs{
				Size:       pulumi.Int(conf.NodePool.DataVolume.Size),
				Volumetype: pulumi.String(conf.NodePool.DataVolume.VolumeType),
			},
		},
		InitialNodeCount: args.NodeCount,
		ClusterId:        clusterRes.ID(),
		SubnetId:         args.SubnetId,
		Name:             pulumi.String(nodepoolName),
		Os:               pulumi.String(conf.NodePool.Os),
		FlavorId:         pulumi.String(conf.NodePool.FlavorId),
		Password:         pulumi.String("Test@1234"), // TODO don't use this in production
	})
	if err != nil {
		return nil, err
	}
	return clusterRes, nil
}
