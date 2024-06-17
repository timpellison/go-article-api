resource "aws_rds_cluster" "rds" {
  cluster_identifier   = "tpe-postgres-1"
  engine               = "aurora-postgresql"
  engine_mode          = "serverless"
  database_name        = "postgres"
  master_username      = "postgres"
  master_password      = random_password.password.result
  db_subnet_group_name = aws_db_subnet_group.rds_subnet_group.name
  availability_zones   = slice(data.aws_availability_zones.available.names, 0, 2)
  skip_final_snapshot  = true

  vpc_security_group_ids      = [aws_security_group.rds_sg.id]
  apply_immediately           = true
  allow_major_version_upgrade = true

  scaling_configuration {
    max_capacity = 16
    min_capacity = 2
  }
}

resource "random_password" "password" {
  length  = 12
  special = false
}

resource "aws_db_subnet_group" "rds_subnet_group" {
  name       = "rds-subnet-group"
  subnet_ids = data.aws_subnets.all.ids
  #  depends_on = [aws_vpc.main]
}

data "aws_subnets" "all" {
  filter {
    name   = "vpc-id"
    values = [aws_default_vpc.default_vpc.id]
    #    values = [aws_vpc.main.id]
  }
}