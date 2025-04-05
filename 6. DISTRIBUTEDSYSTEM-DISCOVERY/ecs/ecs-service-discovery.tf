resource "aws_service_discovery_private_dns_namespace" "micro" {
    name = "micro.local"
    vpc = data.aws_vpc.vpc-main.id
}