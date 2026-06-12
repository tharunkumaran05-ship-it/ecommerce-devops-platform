terraform {
  required_version = ">= 1.0"

  required_providers {
    local = {
      source = "hashicorp/local"
      version = "~> 2.4"
    }
  }
}

resource "local_file" "welcome" {
  filename = "terraform-created.txt"

  content = "Hello ${var.username}! This file was created by Terraform."
}