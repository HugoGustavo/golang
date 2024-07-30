### AWS IAM ACCESS KEY ###
output "iam_access_key_create_date" {
  description = "Date and time in RFC3339 format that the access key was created"
  value       = module.aws_iam_access_key.create_date
}

output "iam_access_key_encrypted_secret" {
  description = "Encrypted secret, base64 encoded, if pgp_key was specified"
  value       = module.aws_iam_access_key.encrypted_secret
  sensitive   = true
}

output "iam_access_key_encrypted_ses_smtp_password_v4" {
  description = "Encrypted SES SMTP password, base64 encoded, if pgp_key was specified"
  value       = module.aws_iam_access_key.encrypted_ses_smtp_password_v4
  sensitive   = true
}

output "iam_access_key_id" {
  description = "Access key ID"
  value       = module.aws_iam_access_key.id
}

output "iam_access_key_key_fingerprint" {
  description = "Fingerprint of the PGP key used to encrypt the secret"
  value       = module.aws_iam_access_key.key_fingerprint
}

output "iam_access_key_secret" {
  description = "Secret access key"
  value       = module.aws_iam_access_key.secret
  sensitive   = true
}

output "iam_access_key_ses_smtp_password_v4" {
  description = "Secret access key converted into an SES SMTP password by applying AWS's documented Sigv4 conversion algorithm"
  value       = module.aws_iam_access_key.ses_smtp_password_v4
  sensitive   = true
}

### AWS IAM GROUP ###
output "iam_group_id" {
  description = "The group's name"
  value       = module.aws_iam_group.id
}

output "iam_group_arn" {
  description = "The ARN assigned by AWS for this user"
  value       = module.aws_iam_group.arn
}

output "iam_group_name" {
  description = "The group's name"
  value       = module.aws_iam_group.name
}

output "iam_group_path" {
  description = "The path of the group in IAM"
  value       = module.aws_iam_group.path
}

output "iam_group_unique_id" {
  description = "The unique ID assigned by AWS"
  value       = module.aws_iam_group.unique_id
}

### AWS IAM GROUP MEMBERSHIP ###
output "iam_group_membership_name" {
  value = module.aws_iam_group_membership.name
}

output "iam_group_membership_users" {
  description = "The name to identify the Group Membership"
  value       = module.aws_iam_group_membership.users
}

output "iam_group_membership_group" {
  description = "IAM Group name"
  value       = module.aws_iam_group_membership.group
}

### AWS IAM USER ###
output "iam_user_arn" {
  description = "The ARN assigned by AWS for this user"
  value       = module.aws_iam_user.arn
}

output "iam_user_id" {
  description = "The user's name"
  value       = module.aws_iam_user.id
}

output "iam_user_name" {
  description = "The user's name"
  value       = module.aws_iam_user.name
}

output "iam_user_tags_all" {
  description = "A map of tags assigned to the resource"
  value       = module.aws_iam_user.tags_all
}

output "iam_user_unique_id" {
  description = "The unique ID assigned by AWS"
  value       = module.aws_iam_user.unique_id
}

### AWS IAM USER LOGIN PROFILE ###
output "iam_user_login_profile_password" {
  description = "The plain text password"
  value       = module.aws_iam_user_login_profile.password
  sensitive   = true
}

output "iam_user_login_profile_key_fingerprint" {
  description = "The fingerprint of the PGP key used to encrypt the password"
  value       = module.aws_iam_user_login_profile.key_fingerprint
}

output "iam_user_login_profile_encrypted_password" {
  description = "The encrypted password, base64 encoded"
  value       = module.aws_iam_user_login_profile.encrypted_password
}

### AWS IAM SSM PARAMETER ###
output "ssm_parameter_arn" {
  description = "ARN of the parameter"
  value       = module.aws_ssm_parameter.arn
}

output "ssm_parameter_tags_all" {
  description = "Map of tags assigned to the resource"
  value       = module.aws_ssm_parameter.tags_all
}

output "ssm_parameter_version" {
  description = "Version of the parameter"
  value       = module.aws_ssm_parameter.version
}

### AWS SECRET MANAGER SECRET ###
output "secretsmanager_secret_id" {
  description = "ARN of the secret"
  value       = module.aws_secretsmanager_secret.id
}

output "secretsmanager_secret_arn" {
  description = "ARN of the secret"
  value       = module.aws_secretsmanager_secret.arn
}

### AWS SECRET MANAGER SECRET VERSION ###
output "secretsmanager_secret_version_arn" {
  description = "ARN of the secret"
  value       = module.aws_secretsmanager_secret_version.arn
}

output "secretsmanager_secret_version_id" {
  description = "The unique identifier of the version of the secret"
  value       = module.aws_secretsmanager_secret_version.id
}

output "secretsmanager_secret_version_secret_string" {
  description = "Specifies text data that you want to encrypt and store in this version of the secret"
  value       = module.aws_secretsmanager_secret_version.secret_string
  sensitive   = true
}

### BUCKET ###
output "s3_bucket_id" {
  description = "Name of the bucket"
  value       = module.aws_s3_bucket.id
  sensitive   = false
}

output "s3_bucket_arn" {
  description = "ARN of the bucket"
  value       = module.aws_s3_bucket.arn
  sensitive   = false
}

output "s3_bucket_region" {
  description = "AWS region this bucket resides in"
  value       = module.aws_s3_bucket.region
  sensitive   = false
}