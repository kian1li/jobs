data "terraform_remote_state" "vpc" {

  backend = "s3"
  config = {
    bucket = format("tf-ecs-state-work-%s", var.tfid)
    region = var.region
    key    = "terraform/terraform_locks_vpc.tfstate"
  }
}