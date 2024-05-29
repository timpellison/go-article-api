resource "aws_ecr_repository" "my_ecr_repository" {
  name                 = "ecr-ue1-article-api-poc"
  image_tag_mutability = "MUTABLE"
  force_delete         = true

  image_scanning_configuration {
    scan_on_push = true

  }
}