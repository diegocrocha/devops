terraform {
  backend "s3" {
    bucket         = "terraform-pc-die"
    key            = "terraform/devops"
    region         = "us-east-1"
  }
}
