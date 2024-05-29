resource "aws_rds_cluster" "rds" {
  cluster_identifier = "tpe-postgres-1"
  engine             = "aurora-postgresql"
  #  availability_zones = ["us-east-1a", "us-east-1b", "us-east-1c"]
  #  vpc_security_group_ids = [module.vpc.default_security_group_id]
  database_name   = "postgres"
  master_username = "postgres"
  master_password = random_password.password.result

  skip_final_snapshot = true

  serverlessv2_scaling_configuration {
    max_capacity = 16
    min_capacity = 2
  }


}

resource "random_password" "password" {
  length  = 12
  special = false
}

