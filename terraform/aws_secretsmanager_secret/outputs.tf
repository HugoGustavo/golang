output "arn" {
  description = "The ARN of the secret"
  value       = aws_secretsmanager_secret.aws_secretsmanager_secret.arn
  sensitive   = false
}

output "id" {
  description = "A pipe delimited combination of secret ID and version ID"
  value       = aws_secretsmanager_secret.aws_secretsmanager_secret.id
  sensitive   = false
}
