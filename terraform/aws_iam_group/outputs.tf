output "id" {
  description = "The group's name"
  value       = aws_iam_group.aws_iam_group.id
  sensitive   = false
}

output "arn" {
  description = "The ARN assigned by AWS for this user"
  value       = aws_iam_group.aws_iam_group.arn
  sensitive   = false
}

output "name" {
  description = "The group's name"
  value       = aws_iam_group.aws_iam_group.name
  sensitive   = false
}

output "path" {
  description = "The path of the group in IAM"
  value       = aws_iam_group.aws_iam_group.path
  sensitive   = false
}

output "unique_id" {
  description = "The unique ID assigned by AWS"
  value       = aws_iam_group.aws_iam_group.unique_id
  sensitive   = false
}