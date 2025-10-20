package main

import (
	"pulumi-quick-start/modules/cluster"
	"pulumi-quick-start/modules/network"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create VPC, subnet, EIP and NAT Gateway
		net, err := network.CreateNetwork(ctx)
		if err != nil {
			return err
		}

		// Create CCE cluster and node pool
		_, err = cluster.CreateCCE(ctx, cluster.CCEArgs{
			VpcId:        net.VpcId,
			SubnetId:     net.SubnetId,
			EipId:        net.EipId,
			NatGatewayId: net.NatGatewayId,
			SubnetCidr:   net.SubnetCidr,
		})
		if err != nil {
			return err
		}

		ctx.Export("natEip", net.EipId)

		return nil
	})
}
