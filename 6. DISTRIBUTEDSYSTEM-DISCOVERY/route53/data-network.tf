data "aws_lb" "router_alb" {
  arn  = data.terraform_remote_state.net.outputs.router_alb.arn
  name = data.terraform_remote_state.net.outputs.router_alb.name
}