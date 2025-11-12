package network

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type NetworkOutput struct {
	VpcId pulumi.IDOutput
	SgID  pulumi.IDOutput
}

type NatOutput struct {
	NatGatewayId pulumi.IDOutput
}

type SnatRuleOutput struct {
	SnatRuleId pulumi.IDOutput
}
