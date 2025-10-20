package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru/nat"
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru/vpc"
)

func CreateNetwork(ctx *pulumi.Context) (*NetworkOutput, error) {
	vpcRes, err := vpc.NewVpc(ctx, "vpc", &vpc.VpcArgs{
		Cidr: pulumi.String("192.168.0.0/16"),
		Name: pulumi.String("my-vpc"),
	})
	if err != nil {
		return nil, err
	}

	subnetRes, err := vpc.NewSubnet(ctx, "subnet", &vpc.SubnetArgs{
		VpcId:     vpcRes.ID(),
		Cidr:      pulumi.String("192.168.10.0/24"),
		GatewayIp: pulumi.String("192.168.10.1"),
		Name:      pulumi.String("my-subnet"),
	})
	if err != nil {
		return nil, err
	}

	eipRes, err := vpc.NewEip(ctx, "eip", &vpc.EipArgs{
		Publicip: &vpc.EipPublicipArgs{
			Type: pulumi.String("5_bgp"),
		},
		Bandwidth: vpc.EipBandwidthArgs{
			Name:       pulumi.String("my-eip-bandwidth"),
			ShareType:  pulumi.String("PER"),
			Size:       pulumi.Int(10),
			ChargeMode: pulumi.String("traffic"),
		},
	})
	if err != nil {
		return nil, err
	}

	natRes, err := nat.NewGateway(ctx, "natgw", &nat.GatewayArgs{

		Spec:     pulumi.String("1"),
		Name:     pulumi.String("my-nat"),
		VpcId:    vpcRes.ID(),
		SubnetId: subnetRes.ID(),
	})
	if err != nil {
		return nil, err
	}

	// SNAT Rule: outbound traffic from cluster's subnet to internet via NAT and EIP
	snatRule, err := nat.NewSnatRule(ctx, "snatRule", &nat.SnatRuleArgs{
		FloatingIpId: eipRes.ID().ToStringOutput(),
		NatGatewayId: natRes.ID(),
		SubnetId:     subnetRes.ID(),
	})
	if err != nil {
		return nil, err
	}

	return &NetworkOutput{
		VpcId:        vpcRes.ID().ToIDOutput(),
		SubnetId:     subnetRes.ID(),
		EipId:        eipRes.ID(),
		NatGatewayId: natRes.ID(),
		SubnetCidr:   subnetRes.Cidr,
		SnatRuleId:   snatRule.ID(),
	}, nil
}
