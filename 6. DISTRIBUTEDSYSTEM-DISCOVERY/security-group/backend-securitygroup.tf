terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
  backend "s3" {
    bucket  = "${var.bucket_id}"
    key     = "terraform/terraform_locks_securitygroup.tfstate"
    region  = "ap-northeast-2"
    encrypt = "true"
  }
}

# Configure the AWS Provider
provider "aws" {
  region = var.region
  default_tags {
    tags = var.default_tags
  }
}
