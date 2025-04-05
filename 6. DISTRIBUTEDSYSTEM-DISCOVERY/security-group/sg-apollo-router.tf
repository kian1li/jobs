# Security Group for apollo_router ALB
resource "aws_security_group" "sg_apollo_router_alb" {
  name_prefix = "${var.default_tags.project}-sg-apollo_router-alb"
  description = "security group for apollo_router service application load balancer"
  vpc_id      = data.aws_vpc.vpc-main.id
}

resource "aws_security_group_rule" "sg_apollo_router_alb_allow_80" {
  security_group_id = aws_security_group.sg_apollo_router_alb.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 80
  to_port           = 80
  cidr_blocks       = ["0.0.0.0/0"]
  ipv6_cidr_blocks  = ["::/0"]
  description       = "Allow HTTP traffic."
}

resource "aws_security_group_rule" "sg_apollo_router_alb_allow_443" {
  security_group_id = aws_security_group.sg_apollo_router_alb.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 443
  to_port           = 443
  cidr_blocks       = ["0.0.0.0/0"]
  ipv6_cidr_blocks  = ["::/0"]
  description       = "Allow HTTP traffic."
}


resource "aws_security_group_rule" "apollo_router_alb_allow_outbound" {
  security_group_id = aws_security_group.sg_apollo_router_alb.id
  type              = "egress"
  protocol          = "-1"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
  ipv6_cidr_blocks  = ["::/0"]
  description       = "Allow any outbound traffic."
}

# Security Group for ECS apollo_router Service.
resource "aws_security_group" "sg_apollo_router_service" {
  name_prefix = "${var.default_tags.project}-sg-apollo_router-service"
  description = "ECS apollo_router service security group."
  vpc_id      = data.aws_vpc.vpc-main.id
}

resource "aws_security_group_rule" "apollo_router_service_allow_9090" {
  security_group_id        = aws_security_group.sg_apollo_router_service.id
  type                     = "ingress"
  protocol                 = "tcp"
  from_port                = var.application_port
  to_port                  = var.application_port
  source_security_group_id = aws_security_group.sg_apollo_router_alb.id
  description              = "Allow incoming traffic from the apollo_router ALB into the service container port."
}

resource "aws_security_group_rule" "apollo_router_service_allow_health_port" {
  security_group_id        = aws_security_group.sg_apollo_router_service.id
  type                     = "ingress"
  protocol                 = "tcp"
  from_port                = var.router_health_port
  to_port                  = var.router_health_port
  source_security_group_id = aws_security_group.sg_apollo_router_alb.id
  description              = "Allow incoming traffic from the apollo_router ALB into the service container port."
}

resource "aws_security_group_rule" "apollo_router_service_allow_inbound_self" {
  security_group_id = aws_security_group.sg_apollo_router_service.id
  type              = "ingress"
  protocol          = -1
  self              = true
  from_port         = 0
  to_port           = 0
  description       = "Allow traffic from resources with this security group."
}

resource "aws_security_group_rule" "apollo_router_service_allow_outbound" {
  security_group_id = aws_security_group.sg_apollo_router_service.id
  type              = "egress"
  protocol          = "-1"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
  ipv6_cidr_blocks  = ["::/0"]
  description       = "Allow any outbound traffic."
}
