resource "aws_iam_access_key" "aws_iam_access_key" {
  user    = var.user
  pgp_key = var.pgp_key
  status  = var.status
}