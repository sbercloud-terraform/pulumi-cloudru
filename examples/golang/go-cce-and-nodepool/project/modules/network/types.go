package network

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type NetworkOutput struct {
	VpcId        pulumi.IDOutput
	SubnetId     pulumi.IDOutput
	EipId        pulumi.IDOutput
	NatGatewayId pulumi.IDOutput
	SubnetCidr   pulumi.StringOutput
	SnatRuleId   pulumi.IDOutput
}
