data "aws_subnets" "private" {
  filter {
    name   = "tag:Name"
    values = ["micro-ecs-subnet-private"]
  }
}

data "aws_lb_target_group" "router_alb_targets" {
  arn  = data.terraform_remote_state.net.outputs.router_alb_targets.arn
  name = data.terraform_remote_state.net.outputs.router_alb_targets.name
}