provider "aws" {
  region = "us-east-1"
}

# provider "docker" {
#   registry_auth {
#     address = module.ecr.repository_url
#     password = data.aws_ecr_authorization_token.this.password
#     username =data.aws_ecr_authorization_token.this.user_name
#   }
# }