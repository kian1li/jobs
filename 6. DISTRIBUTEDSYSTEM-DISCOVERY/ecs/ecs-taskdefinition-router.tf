locals {
  apollo_creds = jsondecode(aws_secretsmanager_secret_version.apollo_cred.secret_string)
}

resource "aws_ecs_task_definition" "router" {
  family                   = "${var.default_tags.project}-router"
  requires_compatibilities = ["FARGATE"]
  memory                   = 512
  cpu                      = 256
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn

  container_definitions = jsonencode([
    {
      name             = "router"
      image            = "${var.user_id}.dkr.ecr.ap-northeast-2.amazonaws.com/apollo-router"
      cpu              = 0
      essential        = true
      task_role_arn    = aws_iam_role.ecs_task_execution_role.arn
      logConfiguration = local.router_logs_configuration
      portMappings = [
        {
          containerPort = var.application_port
          hostPort      = var.application_port
          protocol      = "tcp"
        },
        {
          containerPort = var.router_health_port
          hostPort      = var.router_health_port
          protocol      = "tcp"
        }
      ],
      environment = [
        {
          name  = "NAME"
          value = "router"
        },
        {
          name = "SERVICE_DISCOVERY_NAMESPACE_ID", 
          value = "${aws_service_discovery_private_dns_namespace.micro.id}"
        },
        {
          name  = "APOLLO_KEY"
          value = local.apollo_creds.APOLLO_API_KEY
        },
        {
          name  = "APOLLO_GRAPH_REF"
          value = local.apollo_creds.APOLLO_GRAPH_REF
        }
      ]
    }
  ])
  depends_on = [
    aws_ecs_task_definition.auth,
    aws_ecs_task_definition.file,
    aws_ecs_task_definition.notification,
    aws_ecs_task_definition.post,
    aws_ecs_task_definition.user
  ]
}
