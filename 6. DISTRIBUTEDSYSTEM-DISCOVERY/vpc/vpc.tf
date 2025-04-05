# Create a VPC
resource "aws_vpc" "main" {
  cidr_block = var.vpc_cidr
  tags = {
    "Name" = "${var.default_tags.project}-vpc"
  }
  assign_generated_ipv6_cidr_block = true
  instance_tenancy                 = "default"
  enable_dns_hostnames             = true
  enable_dns_support               = true
}

# resource "aws_cloudwatch_log_group" "vpcflowlog-group" {
#   name              = "vpcflowlogGroup-${var.default_tags.project}"
#   retention_in_days = 30

#   tags = {
#     Name = "vpcflowlog-${var.default_tags.project}"
#   }
# }

# resource "aws_iam_role" "flowlog_role" {
#   name               = "vpcflowlog_role"
#   assume_role_policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement" : [
#     {
#       "Effect": "Allow",
#       "Principal": {
#         "Service": "vpc-flow-logs.amazonaws.com"
#       },
#       "Action": "sts:AssumeRole"
#     }
#   ]
# }
# EOF
#   tags = {
#     Name = "vpcflowlog"
#   }
# }

# resource "aws_iam_role_policy" "flowlog_policy" {
#   name   = "vpcflowlog_policy"
#   role   = aws_iam_role.flowlog_role.id
#   policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement": [
#     {
#       "Action": [
#         "logs:CreateLogGroup",
#         "logs:CreateLogStream",
#         "logs:PutLogEvents",
#         "logs:DescribeLogGroups",
#         "logs:DescribeLogStreams"
#       ],
#       "Effect": "Allow",
#       "Resource": "*"
#     }
#   ]
# }
# EOF
# }

# resource "aws_flow_log" "flowlog" {
#   iam_role_arn    = aws_iam_role.flowlog_role.arn
#   log_destination = aws_cloudwatch_log_group.vpcflowlog-group.arn
#   traffic_type    = "ALL"
#   vpc_id          = aws_vpc.main.id
#   tags = {
#     Name = "vpcflowlog-${var.default_tags.project}"
#   }
# }