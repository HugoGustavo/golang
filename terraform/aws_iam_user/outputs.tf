output "arn" {
  description = "The ARN assigned by AWS for this user"
  value       = aws_iam_user.aws_iam_user.arn
  sensitive   = false
}

output "id" {
  description = "The user's name"
  value       = aws_iam_user.aws_iam_user.id
  sensitive   = false
}

output "name" {
  description = "The user's name"
  value       = aws_iam_user.aws_iam_user.name
  sensitive   = false
}

output "tags_all" {
  description = "A map of tags assigned to the resource"
  value       = aws_iam_user.aws_iam_user.tags_all
  sensitive   = false
}

output "unique_id" {
  description = "The unique ID assigned by AWS"
  value       = aws_iam_user.aws_iam_user.unique_id
  sensitive   = false
}