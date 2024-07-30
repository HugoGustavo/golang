output "create_date" {
  description = "Date and time in RFC3339 format that the access key was created"
  value       = aws_iam_access_key.aws_iam_access_key.create_date
  sensitive   = false
}

output "encrypted_secret" {
  description = "Encrypted secret, base64 encoded, if pgp_key was specified"
  value       = aws_iam_access_key.aws_iam_access_key.encrypted_secret
  sensitive   = false
}

output "encrypted_ses_smtp_password_v4" {
  description = "Encrypted SES SMTP password, base64 encoded, if pgp_key was specified"
  value       = aws_iam_access_key.aws_iam_access_key.encrypted_ses_smtp_password_v4
  sensitive   = false
}

output "id" {
  description = "Access key ID"
  value       = aws_iam_access_key.aws_iam_access_key.id
  sensitive   = false
}

output "key_fingerprint" {
  description = "Fingerprint of the PGP key used to encrypt the secret"
  value       = aws_iam_access_key.aws_iam_access_key.key_fingerprint
  sensitive   = false
}

output "secret" {
  description = "Secret access key"
  value       = aws_iam_access_key.aws_iam_access_key.secret
  sensitive   = true
}

output "ses_smtp_password_v4" {
  description = "Secret access key converted into an SES SMTP password"
  value       = aws_iam_access_key.aws_iam_access_key.ses_smtp_password_v4
  sensitive   = true
}
