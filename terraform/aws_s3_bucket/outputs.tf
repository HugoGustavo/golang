output "id" {
  description = "Name of the bucket"
  value       = aws_s3_bucket.aws_s3_bucket.id
  sensitive   = false
}

output "arn" {
  description = "ARN of the bucket"
  value       = aws_s3_bucket.aws_s3_bucket.arn
  sensitive   = false
}

output "region" {
  description = "AWS region this bucket resides in"
  value       = aws_s3_bucket.aws_s3_bucket.region
  sensitive   = false
}