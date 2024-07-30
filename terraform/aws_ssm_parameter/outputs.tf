output "arn" {
  description = "ARN of the parameter"
  value       = aws_ssm_parameter.aws_ssm_parameter.arn
  sensitive   = false
}

output "tags_all" {
  description = "Map of tags assigned to the resource"
  value       = aws_ssm_parameter.aws_ssm_parameter.tags_all
  sensitive   = false
}

output "version" {
  description = "Version of the parameter"
  value       = aws_ssm_parameter.aws_ssm_parameter.version
  sensitive   = false
}
