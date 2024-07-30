provider "aws" {}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {
    bucket  = "sg-infra-as-code"
    region  = "us-east-1"
    encrypt = true
  }
}