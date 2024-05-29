module "vpc" {
  source                           = "terraform-aws-modules/vpc/aws"
  version                          = "5.8.1"
  default_vpc_enable_dns_hostnames = true
  default_vpc_enable_dns_support   = true
  name                             = "my-vpc"
  cidr                             = var.vpc_cidr

  azs                    = var.vpc_azs
  private_subnets        = var.private_subnets
  public_subnets         = var.public_subnets
  create_igw             = true
  enable_nat_gateway     = true
  create_egress_only_igw = true
  single_nat_gateway     = true

}

resource "aws_vpc_endpoint" "ecr-dkr" {
  service_name = "com.amazonaws.us-east-1.ecr.dkr"
  vpc_id = aws_security_group.alb.vpc_id
  subnet_ids = module.vpc.private_subnets
  vpc_endpoint_type = "Interface"
  security_group_ids = [module.vpc.default_security_group_id]
  private_dns_enabled = true
}

resource "aws_vpc_endpoint" "ecr-api" {
  service_name = "com.amazonaws.us-east-1.ecr.api"
  vpc_id = aws_security_group.alb.vpc_id
  vpc_endpoint_type = "Interface"
  subnet_ids = module.vpc.private_subnets
  security_group_ids = [module.vpc.default_security_group_id]
  private_dns_enabled = true
}


resource "aws_vpc_endpoint" "ecr-s3" {
  service_name = "com.amazonaws.us-east-1.s3"
  vpc_endpoint_type = "Gateway"
  vpc_id = aws_security_group.alb.vpc_id

  route_table_ids = module.vpc.private_route_table_ids

#  private_dns_enabled = true
}

