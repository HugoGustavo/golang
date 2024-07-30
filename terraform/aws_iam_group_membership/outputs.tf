output "name" {
  description = "The name to identify the Group Membership"
  value       = aws_iam_group_membership.aws_iam_group_membership.name
  sensitive   = false
}

output "users" {
  description = "A list of IAM User names"
  value       = aws_iam_group_membership.aws_iam_group_membership.users
  sensitive   = false
}

output "group" {
  description = "IAM Group name"
  value       = aws_iam_group_membership.aws_iam_group_membership.group
  sensitive   = false
}