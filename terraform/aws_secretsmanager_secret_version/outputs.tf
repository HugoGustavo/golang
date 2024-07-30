output "arn" {
  description = "The ARN of the secret"
  value       = aws_secretsmanager_secret_version.aws_secretsmanager_secret_version.arn
  sensitive   = false
}

output "id" {
  description = "A pipe delimited combination of secret ID and version ID"
  value       = aws_secretsmanager_secret_version.aws_secretsmanager_secret_version.id
  sensitive   = false
}

output "secret_string" {
  description = "Specifies text data that you want to encrypt and store in this version of the secret"
  value       = aws_secretsmanager_secret_version.aws_secretsmanager_secret_version.secret_string
  sensitive   = true
}
