data "aws_security_group" "sg_apollo_router_service" {
  id      = data.terraform_remote_state.securitygroup.outputs.sg_apollo_router_service.id
}

data "aws_security_group" "sg_ecs_auth_service" {
  id      = data.terraform_remote_state.securitygroup.outputs.sg_ecs_auth_service.id
}

data "aws_security_group" "sg_ecs_file_service" {
  id      = data.terraform_remote_state.securitygroup.outputs.sg_ecs_file_service.id
}

data "aws_security_group" "sg_ecs_notification_service" {
  id      = data.terraform_remote_state.securitygroup.outputs.sg_ecs_notification_service.id
}

data "aws_security_group" "sg_ecs_post_service" {
  id      = data.terraform_remote_state.securitygroup.outputs.sg_ecs_post_service.id
}

data "aws_security_group" "sg_ecs_user_service" {
  id      = data.terraform_remote_state.securitygroup.outputs.sg_ecs_user_service.id
}