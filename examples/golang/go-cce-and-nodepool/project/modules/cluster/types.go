package cluster

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type CCEArgs struct {
	VpcId        pulumi.IDOutput
	SubnetId     pulumi.IDOutput
	EipId        pulumi.IDInput
	NatGatewayId pulumi.IDInput
	SubnetCidr   pulumi.StringInput
}
