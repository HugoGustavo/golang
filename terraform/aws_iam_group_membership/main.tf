resource "aws_iam_group_membership" "aws_iam_group_membership" {
  name  = var.name
  users = [var.user]
  group = var.group
}