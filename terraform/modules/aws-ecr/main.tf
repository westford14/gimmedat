terraform {
  required_version = ">= 1.5.6"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.19.0"
    }
  }
}

locals {
  actions = [
    "ecr:BatchCheckLayerAvailability",
    "ecr:BatchGetImage",
    "ecr:CompleteLayerUpload",
    "ecr:CreateRepository",
    "ecr:DescribeImages",
    "ecr:DescribeImageScanFindings",
    "ecr:DescribeRepositories",
    "ecr:GetAuthorizationToken",
    "ecr:GetDownloadUrlForLayer",
    "ecr:GetLifecyclePolicy",
    "ecr:GetRepositoryPolicy",
    "ecr:InitiateLayerUpload",
    "ecr:ListImages",
    "ecr:ListTagsForResource",
    "ecr:PutImage",
    "ecr:PutImageScanningConfiguration",
    "ecr:PutImageTagMutability",
    "ecr:StartImageScan",
    "ecr:TagResource",
    "ecr:UntagResource",
    "ecr:UploadLayerPart",
  ]
}

data "aws_iam_policy_document" "organization" {
  statement {
    actions = local.actions
    principals {
      identifiers = ["*"]
      type        = "AWS"
    }
  }
}

resource "aws_ecr_repository" "repository" {
  name = var.repo

  force_delete = true

  tags = {
    Name = var.repo
  }
}

resource "aws_ecr_repository_policy" "policy" {
  policy     = data.aws_iam_policy_document.organization.json
  repository = aws_ecr_repository.repository.name
}

resource "aws_ecr_lifecycle_policy" "lifecycle_policy" {
  repository = aws_ecr_repository.repository.name

  policy = <<EOF
{
  "rules": [
    {
      "rulePriority": 1,
      "description": "Keep last 2500 images",
      "selection": {
        "tagStatus": "any",
        "countType": "imageCountMoreThan",
        "countNumber": 2500
      },
      "action": {
        "type": "expire"
      }
    }
  ]
}
EOF
}