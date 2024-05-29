resource "aws_cloudwatch_log_group" "my_ecs_service_log_group" {
  name = "my-ecs-service-loggroup"

}

resource "aws_cloudwatch_log_stream" "my_ecs_service_log_stream" {
  log_group_name = aws_cloudwatch_log_group.my_ecs_service_log_group.name
  name = "app-logstream"
}