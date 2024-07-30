output "password" {
  description = "The plain text password"
  value       = aws_iam_user_login_profile.aws_iam_user_login_profile.password
  sensitive   = true
}

output "key_fingerprint" {
  description = "The fingerprint of the PGP key used to encrypt the password"
  value       = aws_iam_user_login_profile.aws_iam_user_login_profile.key_fingerprint
  sensitive   = false
}

output "encrypted_password" {
  description = "The encrypted password, base64 encoded"
  value       = aws_iam_user_login_profile.aws_iam_user_login_profile.encrypted_password
  sensitive   = false
}