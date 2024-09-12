provider "aws" {
  region = var.region
  default_tags {
    tags = {
      env       = var.environment
      region    = var.region
      terraform = "true"
      project   = var.project
    }
  }
}

terraform {
  required_version = "~> 1.5.6"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.19.0"
    }
  }

  backend "s3" {
    bucket = "warnings-terraform"
    key    = "state/prod.tfstate"
    region = "eu-north-1"
  }
}

module "ecr" {
  source = "../../modules/aws-ecr"
  repo   = var.ecr_repo
}

module "bucket" {
  source = "../../modules/aws-s3"
  bucket = var.bucket_name
}
