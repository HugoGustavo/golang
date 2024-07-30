resource "aws_s3_bucket" "aws_s3_bucket" {
  bucket        = substr(var.bucket, 0, 62)
  force_destroy = var.force_destroy
}