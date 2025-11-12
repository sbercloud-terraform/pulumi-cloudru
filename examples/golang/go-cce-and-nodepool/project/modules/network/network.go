package network

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru/nat"
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru/vpc"
)

func CreateSubnet(ctx *pulumi.Context, vpcId pulumi.IDOutput, cidr, name string) (*vpc.Subnet, error) {
	return vpc.NewSubnet(ctx, name, &vpc.SubnetArgs{
		VpcId:     vpcId,
		Cidr:      pulumi.String(cidr),
		GatewayIp: pulumi.String(strings.Replace(cidr, "0/24", "1", 1)),
		Name:      pulumi.String(name),
	})
}

func CreateNetwork(ctx *pulumi.Context) (*NetworkOutput, error) {
	vpcRes, err := vpc.NewVpc(ctx, "main-vpc", &vpc.VpcArgs{
		Cidr: pulumi.String("192.168.0.0/16"),
		Name: pulumi.String("main-vpc"),
	})
	if err != nil {
		return nil, err
	}

	sgRes, err := vpc.NewSecgroup(ctx, "mainSG", &vpc.SecgroupArgs{
		Name:        pulumi.String("main-sg"),
		Description: pulumi.String("Security group for cluster ingress"),
	})
	if err != nil {
		return nil, err
	}

	return &NetworkOutput{
		VpcId: vpcRes.ID(),
		SgID:  sgRes.ID(),
	}, nil
}

func CreateEip(ctx *pulumi.Context, name string) (*vpc.Eip, error) {
	eipRes, err := vpc.NewEip(ctx, name+"-eip", &vpc.EipArgs{
		Publicip: &vpc.EipPublicipArgs{
			Type: pulumi.String("5_bgp"),
		},
		Bandwidth: vpc.EipBandwidthArgs{
			Name:       pulumi.String(name + "eip-bandwidth"),
			ShareType:  pulumi.String("PER"),
			Size:       pulumi.Int(10),
			ChargeMode: pulumi.String("traffic"),
		},
	})
	if err != nil {
		return nil, err
	}

	return eipRes, nil
}

func CreateNatGateway(ctx *pulumi.Context, vpcID, subnetID pulumi.IDOutput, name string) (*NatOutput, error) {
	gwName := name + "-natgw"

	natRes, err := nat.NewGateway(ctx, gwName, &nat.GatewayArgs{
		Spec:     pulumi.String("1"),
		Name:     pulumi.String(gwName),
		VpcId:    vpcID,
		SubnetId: subnetID,
	})
	if err != nil {
		return nil, err
	}

	return &NatOutput{
		NatGatewayId: natRes.ID(),
	}, nil
}

func CreateSnatRule(ctx *pulumi.Context, name string, eipID, natID, subnetID pulumi.IDOutput) (*SnatRuleOutput, error) {
	// SNAT Rule: outbound traffic from cluster's subnet to internet via NAT and EIP
	snatRule, err := nat.NewSnatRule(ctx, name+"-snatRule", &nat.SnatRuleArgs{
		FloatingIpId: eipID,
		NatGatewayId: natID,
		SubnetId:     subnetID,
	})
	if err != nil {
		return nil, err
	}

	return &SnatRuleOutput{SnatRuleId: snatRule.ID()}, nil
}

func CreateSgRule(ctx *pulumi.Context, sgId pulumi.IDOutput, sourceIp string, name string) (*vpc.SecgroupRule, error) {
	rule, err := vpc.NewSecgroupRule(ctx, name, &vpc.SecgroupRuleArgs{
		Direction:       pulumi.String("ingress"),
		SecurityGroupId: sgId,
		RemoteIpPrefix:  pulumi.String(sourceIp),
		Ethertype:       pulumi.String("IPv4"),
		Protocol:        pulumi.String("tcp"),
		Ports:           pulumi.String("443"),
		Action:          pulumi.String("allow"),
		Description:     pulumi.String(fmt.Sprintf("Allow TCP from %s", sourceIp)),
	})
	return rule, err
}
