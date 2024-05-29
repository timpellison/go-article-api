data "aws_caller_identity" "this" {

}

data "aws_region" "this" {

}

locals {
  app_name       = replace(lower(var.app_name), " ", "-")
  env_name       = replace(lower(var.env_name), " ", "-")
  ecr_address    = format("%v.dkr.ecr.%v.amazonaws.com", data.aws_caller_identity.this.account_id, data.aws_region.this.name)
  cluster_name   = "ue1-${local.env_name}-${local.app_name}"
  container_port = 8080
}

