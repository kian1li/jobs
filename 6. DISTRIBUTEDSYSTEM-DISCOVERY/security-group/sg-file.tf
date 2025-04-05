# Security Group for ECS File Service.
resource "aws_security_group" "sg_ecs_file_service" {
  name_prefix = "${var.default_tags.project}-sg-file-service"
  description = "ECS File service security group."
  vpc_id      = data.aws_vpc.vpc-main.id
}

resource "aws_security_group_rule" "ecs_file_service_allow_same_vpc" {
  security_group_id = aws_security_group.sg_ecs_file_service.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = var.application_port
  to_port           = var.application_port
  cidr_blocks       = [var.vpc_cidr]
  description       = "Allow incoming traffic from the into same VPC the service notification container port."
}

resource "aws_security_group_rule" "ecs_file_service_allow_grpc" {
  security_group_id = aws_security_group.sg_ecs_file_service.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = var.file_grpc_port
  to_port           = var.file_grpc_port
  cidr_blocks       = [var.vpc_cidr]
  description       = "Allow incoming traffic from the file ALB into the service container grpc port."
}

resource "aws_security_group_rule" "ecs_file_service_allow_inbound_self" {
  security_group_id = aws_security_group.sg_ecs_file_service.id
  type              = "ingress"
  protocol          = -1
  self              = true
  from_port         = 0
  to_port           = 0
  description       = "Allow traffic from resources with this security group."
}

resource "aws_security_group_rule" "ecs_file_service_allow_outbound" {
  security_group_id = aws_security_group.sg_ecs_file_service.id
  type              = "egress"
  protocol          = "-1"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
  ipv6_cidr_blocks  = ["::/0"]
  description       = "Allow any outbound traffic."
}
