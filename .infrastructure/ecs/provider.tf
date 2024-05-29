terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }

    #     docker = {
    #       source  = "kreuzwerker/docker"
    #       version = "~> 3.0"
    #     }
  }

  backend "s3" {
    bucket = "tpe-ue1-terraform-state-bucket"
    key    = "state/terraform_state.tfstate"
    region = "us-east-1"
  }
}