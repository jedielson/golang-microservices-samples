terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    rabbitmq = {
      source = "cyrilgdn/rabbitmq"
    }
  }
  required_version = ">= 0.13"
}
