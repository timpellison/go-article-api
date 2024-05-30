resource "aws_ecr_repository" "goservice-repo" {
  name = var.ecr-repo-name
  image_scanning_configuration {
    scan_on_push = true
  }
  force_delete = true
}