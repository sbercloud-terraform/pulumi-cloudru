## Quick start guide 
Pulumi is an infrastructure-as-code tool similar to Terraform. On this page, you’ll find instructions for deploying a VPC in the advanced cloud. More detailed instructions and a description of the pros/cons will be available in other articles.

This version is for **Mac OS** and the **Go** language; the installation process for Pulumi and the SDK may differ on other systems and languages.

Install Pulumi with the following command (if you have Homebrew installed):
```shell
brew install pulumi/tap/pulumi
```

Or with this command:
```shell
curl -fsSL https://get.pulumi.com | sh
```

Then execute 
```shell
pulumi version
```

If everything is installed correctly, you should see output similar to:
```shell
v3.202.0
```

Next, open your favorite code editor (VS Code is used in this example) and navigate to your project directory (create a new empty one if needed).
Open a terminal in this directory and run:
```shell
pulumi new <language>
```

For this example run:
```shell
pulumi new go
```

After this, enter all the required details prompted by the program to initialize your project.
Alternatively, you can run everything as a single command (replace with your own values):
```shell
# In the first line, you set a password; it can be left empty
export PULUMI_CONFIG_PASSPHRASE='my-strong-password'
pulumi new go \
  --name=pulumi-quick-start \
  --description="Demo VPC in CloudRu" \
  --stack=dev \
  --yes
```

After execution, the files go.mod, go.sum, main.go, and Pulumi.yaml will be created.

Install SDK with:
```shell
go get github.com/sbercloud-terraform/pulumi-cloudru/sdk
```

You also need to install the plugin. Since the plugin is not yet published in the official Pulumi registry, 
you need to explicitly specify where and which version to download.
To find the version, visit https://github.com/sbercloud-terraform/pulumi-cloudru/releases and choose 
the required one. In the example below, it is version `v1.0.0-alpha.1760368627`.
```shell
pulumi plugin install resource cloudru v1.0.0-alpha.1760458278 \
--server "https://github.com/sbercloud-terraform/pulumi-cloudru/releases/download/v1.0.0-alpha.1760458278/"
```
Create an AK/SK pair for programmatic access if you don’t already have one.
Instructions: https://cloud.ru/docs/obs/ug/topics/guides__create-access-keys
Set pulumi variables:
```shell
pulumi config set cloudru:region ru-moscow-1
pulumi config set cloudru:access_key <your_access_key>
pulumi config set cloudru:secret_key <your_secret_key> --secret
```

Now you can write your first program. As an example, let's create a VPC.
Below is the content of the main.go file.

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	// required imports for working with the cloudru provider
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru"
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru/vpc"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// VPC description
		vpc.NewVpc(ctx, "test-custom-pulumi", &vpc.VpcArgs{
			Name: pulumi.StringPtr("test-custom-pulumi"),
			Cidr: pulumi.String("192.168.0.0/16"),
		})

		return nil
	})
}
```

After writing the code, run in the console:

```shell
pulumi up
```

You’ll see output similar to:

```shell
Previewing update (dev):
     Type                         Name                    Plan       
 +   pulumi:pulumi:Stack          pulumi-quick-start-dev  create     
 +   └─ cloudru:Vpc:Vpc           test-custom-pulumi      create     

Resources:
    + 2 to create

Do you want to perform this update?  [Use arrows to move, type to filter]
  yes
> no
  details
```
Select yes and press Enter. Then you’ll see:

```shell
Updating (dev):
     Type                         Name                    Status                
 +   pulumi:pulumi:Stack          pulumi-quick-start-dev  created (6s)          
 +   └─ cloudru:Vpc:Vpc           test-custom-pulumi      created (6s)          

Resources:
    + 2 created
Duration: 11s
```

**The VPC has been created in the cloud, you have achieved the desired result!**

If you need to use more than one pulumi provider, you need to initialize each provider explicitly
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	// required imports for working with the cloudru provider
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru"
	"github.com/sbercloud-terraform/pulumi-cloudru/sdk/go/cloudru/vpc"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Or use any other way to provide credentials
		ak := os.Getenv("AK")
		sk := os.Getenv("SK")

		// Initialize the provider
		provider, _ := cloudru.NewProvider(ctx, "ru-moscow-1", &cloudru.ProviderArgs{
			Region: pulumi.StringPtr("ru-moscow-1"),

			AccessKey: pulumi.String(ak),
			SecretKey: pulumi.String(sk),
		})

		// VPC description
		vpc.NewVpc(ctx, "test-custom-pulumi", &vpc.VpcArgs{
			Name: pulumi.StringPtr("test-custom-pulumi"),
			Cidr: pulumi.String("192.168.0.0/16"),
		}, pulumi.Provider(provider)) // pass target provider as an additional argument

		return nil
	})
}
```