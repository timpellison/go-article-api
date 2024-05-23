resource "aws_rds_cluster" "rds" {
  cluster_identifier = "tpe-postgres-1"
  engine             = "aurora-postgresql"
  availability_zones = ["us-east-1a", "us-east-1b", "us-east-1c"]
  database_name      = "postgres"
  master_username    = "postgres"
  master_password    = random_password.password.result
  serverlessv2_scaling_configuration {
    max_capacity = 2
    min_capacity = 16
  }
}

resource "random_password" "password" {
  length = 12
}

