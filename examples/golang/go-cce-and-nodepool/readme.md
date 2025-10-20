## CCE cluster and node pool example
\
In this example we will create following resources: 
- VPC and subnet
- EIP 
- NAT gateway and SNAT rule 
- CCE cluster
- Node pool

Don't forget to install dependencies and configure provider,
using our quick start guide.

### Project structure 

- **main.go** — Entry point; orchestrates the infrastructure build.
- **modules/network** — Provisioning of VPC, Subnet, EIP, NAT, and SNAT Rule.
- **modules/cluster** — Creation of the CCE cluster and Node Pool.

### Important notes
- The node pool password is present in plaintext in this example, 
don't use this in production.
- Clean up all resources when done using:
```shell
pulumi destroy
```

### How it works
- `network.CreateNetwork`: provisions and wires together the VPC, subnet, EIP, NAT Gateway, 
and SNAT rule, returning their IDs.
- `cluster.CreateCCE`: receives the network outputs and creates a CCE cluster with a Node Pool of two worker nodes. 
