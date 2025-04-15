terraform {
    required_providers {
        aws = {
        source  = "hashicorp/aws"
        version = "~> 4.0"
        }
    }
    
    required_version = ">= 1.0.0"
}

provider "aws" {
    region = var.region
    profile = "TiagoMoreira003"
}

resource "aws_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "main" {
  vpc_id     = aws_vpc.vpc.id
  cidr_block = "10.0.0.0/19"

  tags = {
    Name = "Main"
  }
}