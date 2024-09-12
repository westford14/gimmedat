terraform {
  required_version = ">= 1.5.6"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.19.0"
    }
  }
}

resource "aws_s3_bucket_versioning" "versioning" {
  count  = var.versioning ? 1 : 0
  bucket = var.bucket
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket" "default" {
  bucket        = var.bucket
  force_destroy = var.force_destroy
}

resource "aws_s3_bucket_public_access_block" "default" {
  bucket                  = aws_s3_bucket.default.id
  block_public_acls       = var.block_public_acls
  block_public_policy     = true
  restrict_public_buckets = true
  ignore_public_acls      = true
  depends_on = [
    aws_s3_bucket.default
  ]
}