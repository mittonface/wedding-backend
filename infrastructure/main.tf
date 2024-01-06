terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket = "wedding-website-state-bucket"
    key    = "terraform_state.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region     = var.aws_region
}