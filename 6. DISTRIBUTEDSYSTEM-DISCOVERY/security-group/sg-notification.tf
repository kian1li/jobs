# Security Group for ECS notification Service.
resource "aws_security_group" "sg_ecs_notification_service" {
  name_prefix = "${var.default_tags.project}-sg-notification-service"
  description = "ECS notification service security group."
  vpc_id      = data.aws_vpc.vpc-main.id
}

resource "aws_security_group_rule" "ecs_notification_service_allow_same_vpc" {
  security_group_id = aws_security_group.sg_ecs_notification_service.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = var.application_port
  to_port           = var.application_port
  cidr_blocks       = [var.vpc_cidr]
  description       = "Allow incoming traffic from the into same VPC the service notification container port."
}

resource "aws_security_group_rule" "ecs_notification_service_allow_grpc" {
  security_group_id = aws_security_group.sg_ecs_notification_service.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = var.notification_grpc_port
  to_port           = var.notification_grpc_port
  cidr_blocks       = [var.vpc_cidr]
  description       = "Allow incoming traffic from the notification ALB into the service container grpc port."
}

resource "aws_security_group_rule" "ecs_notification_service_allow_inbound" {
  security_group_id = aws_security_group.sg_ecs_notification_service.id
  type              = "ingress"
  protocol          = -1
  self              = true
  from_port         = 0
  to_port           = 0
  description       = "Allow traffic from resources with this security group."
}

resource "aws_security_group_rule" "ecs_notification_service_allow_outbound" {
  security_group_id = aws_security_group.sg_ecs_notification_service.id
  type              = "egress"
  protocol          = "-1"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
  ipv6_cidr_blocks  = ["::/0"]
  description       = "Allow any outbound traffic."
}
