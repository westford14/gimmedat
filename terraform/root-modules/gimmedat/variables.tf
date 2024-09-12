# Common
variable "region" {
  type    = string
  default = "eu-north-1"
}

variable "project" {
  type    = string
  default = "warnings"
}

variable "environment" {
  type = string
}

# ECR
variable "ecr_repo" {
  type = string
}

# S3
variable "bucket_name" {
  type = string
}
