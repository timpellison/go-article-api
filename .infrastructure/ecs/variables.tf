variable "region" {
  type    = string
  default = "us-east-1"
}

variable "vpc_cidr" {
  type        = string
  description = "VPC cidr"
  default     = "10.0.0.0/16"
}

variable "vpc_azs" {
  type        = list(string)
  description = "Availability zones for VPC"
  default     = ["us-east-1a", "us-east-1b", "us-east-1c"]
}

variable "private_subnets" {
  type        = list(string)
  description = "Private subnets inside the VPC"
  default     = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
}

variable "public_subnets" {
  type        = list(string)
  description = "Public subnets inside the VPC"
  default     = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]
}

variable "app_name" {
  type    = string
  default = "article-api"
}

variable "env_name" {
  type    = string
  default = "pod"
}

variable "container_port" {
  type    = number
  default = 8080
}