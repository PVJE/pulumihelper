package pulumihelper

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMapKeys(strings map[string]string) []string {
	keys := make([]string, 0, len(strings))

	for k := range strings {
		keys = append(keys, k)
	}
	return keys
}

func GetMapValues(strings map[string]string) []string {

	values := make([]string, 0, len(strings))

	for _, v := range strings {
		values = append(values, v)
	}
	return values
}

func ToPulumiStringMap(tags map[string]string) pulumi.StringMap {
	tagsMap := make(pulumi.StringMap)
	for k, v := range tags {
		tagsMap[k] = pulumi.String(v)
	}
	return tagsMap
}

func CreateIgw(ctx *pulumi.Context, name string, allocation_id string, subnet_id string, connectivity_type string, tagsMap pulumi.StringMap) (*ec2.NatGateway, error) {
	if len(allocation_id) > 0 && len(connectivity_type) > 0 {
		my_igw, err := ec2.NewNatGateway(ctx, name, &ec2.NatGatewayArgs{
			AllocationId:     pulumi.String(allocation_id),
			ConnectivityType: pulumi.String(connectivity_type),
			SubnetId:         pulumi.String(subnet_id),
			Tags:             pulumi.StringMap(tagsMap),
		})
		if err != nil {
			return nil, err
		}
		return my_igw, err
	} else {
		my_igw, err := ec2.NewNatGateway(ctx, name, &ec2.NatGatewayArgs{
			SubnetId: pulumi.String(subnet_id),
			Tags:     pulumi.StringMap(tagsMap),
		})
		if err != nil {
			return nil, err
		}
		return my_igw, err
	}

}

func CreateVPC(vpc_name string, vpc_assign_generated_ipv6_cidr_block bool, vpc_cidr string, vpc_enable_dns_support bool, vpc_dns_hostanme bool, vpc_instance_tenancy string, tagsMap pulumi.StringMap, ctx *pulumi.Context) (*ec2.Vpc, error) {
	my_vpc, err := ec2.NewVpc(ctx, string(vpc_name), &ec2.VpcArgs{
		AssignGeneratedIpv6CidrBlock: pulumi.Bool(bool(vpc_assign_generated_ipv6_cidr_block)),
		CidrBlock:                    pulumi.String(string(vpc_cidr)),
		EnableDnsSupport:             pulumi.Bool(bool(vpc_enable_dns_support)),
		EnableDnsHostnames:           pulumi.Bool(bool(vpc_dns_hostanme)),
		InstanceTenancy:              pulumi.String(string(vpc_instance_tenancy)),
		Tags:                         pulumi.StringMap(tagsMap),
	}, pulumi.Protect(false))
	if err != nil {
		return nil, err
	}
	return my_vpc, nil
}

func CreateSubnet(vpc_id *ec2.Vpc, name string, assignIpv6AddressOnCreation bool, availabilityZoneId string,
	cidr_block string, availabilityZone string, customerOwnedIpv4Pool string, mapCustomerOwnedIpOnLaunch bool,
	MapPublicIpOnLaunch bool, Outpostarn string,
	tagsMap pulumi.StringMap, ctx *pulumi.Context) (*ec2.Subnet, error) {
	if len(availabilityZone) > 0 {
		my_subnet, err := ec2.NewSubnet(ctx, string(name), &ec2.SubnetArgs{
			VpcId:                       vpc_id.ID(),
			AssignIpv6AddressOnCreation: pulumi.Bool(assignIpv6AddressOnCreation),
			AvailabilityZoneId:          pulumi.String(availabilityZoneId),
			CustomerOwnedIpv4Pool:       pulumi.String(customerOwnedIpv4Pool),
			CidrBlock:                   pulumi.String(string(cidr_block)),
			AvailabilityZone:            pulumi.String(string(availabilityZone)),
			MapCustomerOwnedIpOnLaunch:  pulumi.Bool(bool(mapCustomerOwnedIpOnLaunch)),
			MapPublicIpOnLaunch:         pulumi.Bool(bool(MapPublicIpOnLaunch)),
			OutpostArn:                  pulumi.String(Outpostarn),
			Tags:                        pulumi.StringMap(tagsMap),
		})
		if err != nil {
			return nil, err
		}
		return my_subnet, err
	} else {
		my_subnet, err := ec2.NewSubnet(ctx, string(name), &ec2.SubnetArgs{
			VpcId:                       vpc_id.ID(),
			AssignIpv6AddressOnCreation: pulumi.Bool(assignIpv6AddressOnCreation),
			//AvailabilityZoneId:          pulumi.String(availabilityZoneId),
			CustomerOwnedIpv4Pool:      pulumi.String(customerOwnedIpv4Pool),
			CidrBlock:                  pulumi.String(string(cidr_block)),
			AvailabilityZone:           pulumi.String(string(availabilityZone)),
			MapCustomerOwnedIpOnLaunch: pulumi.Bool(bool(mapCustomerOwnedIpOnLaunch)),
			MapPublicIpOnLaunch:        pulumi.Bool(bool(MapPublicIpOnLaunch)),
			OutpostArn:                 pulumi.String(Outpostarn),
			Tags:                       pulumi.StringMap(tagsMap),
		})
		if err != nil {
			return nil, err
		}
		return my_subnet, err
	}
}
