resource "aws_ecs_cluster" "my_ecs_cluster" {
  name = "my-ecs-cluster"
}

resource "aws_ecs_service" "my_ecs_service" {
  name                               = "my-ecs-service"
  cluster                            = aws_ecs_cluster.my_ecs_cluster.id
  task_definition                    = aws_ecs_task_definition.main_task_definition.arn
  desired_count                      = 1
  deployment_minimum_healthy_percent = 50
  deployment_maximum_percent         = 200
  launch_type                        = "FARGATE"
  scheduling_strategy                = "REPLICA"
  force_new_deployment               = true
  platform_version                   = "1.3.0"


  network_configuration {
    security_groups = [aws_security_group.ecs_services.id]
    subnets         = module.vpc.private_subnets
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.my_ecs_target_group.arn
    container_name   = var.app_name
    container_port   = var.container_port
  }
}

resource "aws_ecs_task_definition" "main_task_definition" {
  family                   = "my-task-definition"
  requires_compatibilities = ["FARGATE"]


  execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn      = aws_iam_role.ecs_task_role.arn
  network_mode       = "awsvpc"
  cpu                = 256
  memory             = 512



  container_definitions = templatefile("containers/task.tpl.json",
    { CONTAINER_PORT    = var.container_port,
      REGION            = var.region,
      LOG_GROUP         = aws_cloudwatch_log_group.my_ecs_service_log_group.name,
      APP_NAME          = var.app_name,
      DATABASE_CLUSTER  = aws_rds_cluster.rds.endpoint,
      DATABASE_USERNAME = aws_rds_cluster.rds.master_username,
      DATABASE_PASSWORD = random_password.password.result
  })

}