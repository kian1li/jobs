output "vpc_id" {
  value       = aws_vpc.main.id
  description = "VPC name of the ECS"
}

output "cidr_block" {
  value       = aws_vpc.main.cidr_block
  description = "VPC cidr block"
}

output "ipv6_cidr_block" {
  value       = aws_vpc.main.ipv6_cidr_block
  description = "VPC ipv6 cidr block"
}

#  cidr_block                      = cidrsubnet(aws_vpc.main.cidr_block, 4, count.index)
#  ipv6_cidr_block                 = cidrsubnet(aws_vpc.main.ipv6_cidr_block, 8, count.index)