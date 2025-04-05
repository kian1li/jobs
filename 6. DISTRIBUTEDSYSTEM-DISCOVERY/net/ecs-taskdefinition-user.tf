resource "aws_ecs_task_definition" "user" {
  family                   = "${var.default_tags.project}-user"
  requires_compatibilities = ["FARGATE"]
  memory                   = 512
  cpu                      = 256
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  
  runtime_platform {
    cpu_architecture = "ARM64"
  }

  container_definitions = jsonencode([
    {
      name          = "user"
      image         = "${var.user_id}.dkr.ecr.ap-northeast-2.amazonaws.com/user-app"
      cpu           = 0
      essential     = true
      task_role_arn = aws_iam_role.ecs_task_execution_role.arn
      logConfiguration = local.user_logs_configuration
      portMappings = [
        {
          containerPort = var.application_port
          hostPort      = var.application_port
          protocol      = "tcp"
        },
        {
          containerPort = var.user_grpc_port
          hostPort      = var.user_grpc_port
          protocol      = "tcp"
        }
      ],
      environment = [
        {
          name  = "NAME"
          value = "user"
        },
        {
          name = "SERVICE_DISCOVERY_NAMESPACE_ID", 
          value = "${aws_service_discovery_private_dns_namespace.micro.id}"
        },
        {
          name  = "USER_APP_HOST"
          value = "${var.application_host}"
        },
        {
          name  = "USER_APP_PORT"
          value = "${tostring(var.application_port)}"
        },
        {
          name = "USER_GRPC_PORT"
          value = "${tostring(var.user_grpc_port)}"
        },
        {
          name  = "AUTH_APP_HOST"
          value = "${aws_service_discovery_service.auth.name}.${aws_service_discovery_private_dns_namespace.micro.name}"
        },
        {
          name  = "AUTH_GRPC_PORT"
          value = "${tostring(var.auth_grpc_port)}"
        },
        {
          name  = "FILE_APP_HOST"
          value = "${aws_service_discovery_service.file.name}.${aws_service_discovery_private_dns_namespace.micro.name}"
        },
        {
          name  = "FILE_GRPC_PORT"
          value = "${tostring(var.file_grpc_port)}"
        },
        {
          name  = "NOTIFICATION_APP_HOST"
          value = "${aws_service_discovery_service.notification.name}.${aws_service_discovery_private_dns_namespace.micro.name}"
        },
        {
          name  = "NOTIFICATION_GRPC_PORT"
          value = "${tostring(var.notification_grpc_port)}"
        },
        {
          name  = "POST_APP_HOST"
          value = "${aws_service_discovery_service.post.name}.${aws_service_discovery_private_dns_namespace.micro.name}"
        },
        {
          name  = "POST_GRPC_PORT"
          value = "${tostring(var.post_grpc_port)}"
        },
      ]
    }
  ])
}