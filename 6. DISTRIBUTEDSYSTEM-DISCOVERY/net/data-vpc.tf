data "aws_vpc" "vpc-main" {
  default = false
  id      = data.terraform_remote_state.vpc.outputs.vpc_id
}