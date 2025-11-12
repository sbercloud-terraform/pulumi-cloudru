## CCE Cluster and Node Pool Example

This example demonstrates the capability to use Go language tools (for reading configuration from JSON and performing HTTP requests) during the infrastructure deployment process, which is not possible with Terraform.

Important note: The configurations for EIP, VPC, and subnet are hardcoded. JSON parsing configuration is only applied for CCE clusters.

The example creates the following resources:
- VPC
- Security Group
    - Ingress rules are created within the SG to allow incoming connections from specific IP addresses
    - Simulates fetching IP addresses for the rules (for the previous point) from an API
- 2 Cloud Container Engine clusters
    - Node pool in each cluster
- 2 EIPs and 2 bandwidths
- 2 NAT Gateways (1 NAT GW per cluster)
    - 2 SNAT rules are created in each NAT GW for the subnets where the CCE clusters and node pool nodes are located
    - Previously created EIPs are used for SNAT

Project structure:
```shell
├── clusters.json      # Cluster configuration description
├── main.go 
├── modules
│   ├── cluster
│   │   ├── cluster.go # Cluster-related functions
│   │   └── types.go   # Custom structures
│   ├── network
│   │   ├── network.go # Network infrastructure-related functions
│   │   └── types.go   # Custom structures
│   └── tech           # Functions for fetching IP lists from API
│       └── allowed_ips.go 
└── readme.md
```

Don't forget to install the necessary dependencies and configure the provider according to the quick start guide.

Do not use the example password for virtual machines in the node pool.