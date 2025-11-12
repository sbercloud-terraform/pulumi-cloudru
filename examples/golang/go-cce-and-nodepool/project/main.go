package main

import (
	"encoding/json"
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
	"pulumi-quick-start/modules/cluster"
	"pulumi-quick-start/modules/network"
	"pulumi-quick-start/modules/tech"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Create VPC and security group
		// Constants for network are hardcoded inside each function,
		// it will be better to make special config for this
		net, err := network.CreateNetwork(ctx)
		if err != nil {
			return err
		}

		// Get allowed ips list from server (dummy inside)
		ips, _ := tech.FetchIpAllowList("https://demo/api/ip-allowlist")

		// Create ingress rules for each IP on port 443
		for i, ip := range ips {
			_, err := network.CreateSgRule(
				ctx,
				net.SgID,
				ip,
				fmt.Sprintf("allow-ip-%d", i),
			)
			if err != nil {
				return err
			}
		}

		// Read external cluster config from JSON at runtime (you can't do this using terraform)
		raw, err := os.ReadFile("clusters.json")
		if err != nil {
			return err
		}
		var clusters []cluster.ClusterConf

		err = json.Unmarshal(raw, &clusters)
		if err != nil {
			return err
		}

		// Create CCE clusters using config
		for _, c := range clusters {
			subnet, err := network.CreateSubnet(ctx, net.VpcId, c.SubnetCidr, c.Name+"-subnet")
			if err != nil {
				return err
			}

			eip, err := network.CreateEip(ctx, c.Name)
			if err != nil {
				return err
			}

			nat, err := network.CreateNatGateway(ctx, net.VpcId, subnet.ID(), c.Name)
			if err != nil {
				return err
			}

			_, err = network.CreateSnatRule(ctx, c.Name, eip.ID(), nat.NatGatewayId, subnet.ID())
			if err != nil {
				return err
			}

			// Create cluster according to JSON config
			_, err = cluster.CreateCCE(ctx, cluster.CCEArgs{
				Name:       pulumi.String(c.Name),
				VpcId:      net.VpcId,
				SubnetId:   subnet.ID(),
				SubnetCidr: subnet.Cidr,
				NodeCount:  pulumi.Int(c.PoolSize),
			}, &c)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
