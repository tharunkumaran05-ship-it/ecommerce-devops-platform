resource "local_file" "project_info" {
  filename = "${path.module}/terraform-created.txt"

  content = <<EOT
Project: Ecommerce DevOps Platform

Technologies:
- Docker
- Kubernetes
- Helm
- Terraform

Created by Terraform
EOT
}