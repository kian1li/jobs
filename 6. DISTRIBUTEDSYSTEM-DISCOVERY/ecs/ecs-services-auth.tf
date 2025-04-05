resource "aws_service_discovery_service" "auth" {
  name = "auth"
  dns_config {
    namespace_id = aws_service_discovery_private_dns_namespace.micro.id
    dns_records {
      ttl  = 10
      type = "A"
    }
    dns_records {
      ttl  = 10
      type = "SRV"
    }
    routing_policy = "MULTIVALUE"
  }
  health_check_custom_config {
    failure_threshold = 1
  }
}

resource "aws_ecs_service" "auth" {
  name            = "${var.default_tags.project}-auth"
  cluster         = aws_ecs_cluster.main.arn
  task_definition = aws_ecs_task_definition.auth.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = data.aws_subnets.private.ids
    assign_public_ip = false
    security_groups  = [data.aws_security_group.sg_ecs_auth_service.id]
  }

  service_registries {
    registry_arn   = "${aws_service_discovery_service.auth.arn}"
    container_name = "auth"
    container_port = var.application_port
  }
}
