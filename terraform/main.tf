module "aws_iam_access_key" {
  source  = "./aws_iam_access_key"
  user    = var.user
  pgp_key = var.pgp_key
  status  = var.status
}

module "aws_iam_group" {
  source = "./aws_iam_group"
  name   = var.name
}

module "aws_iam_group_membership" {
  source = "./aws_iam_group_membership"
  name   = var.name
  user   = var.user
  group  = var.group
}

module "aws_iam_user" {
  source               = "./aws_iam_user"
  name                 = var.name
  permissions_boundary = var.permissions_boundary
  force_destroy        = var.force_destroy
  tags                 = var.tags
}

module "aws_iam_user_login_profile" {
  source                  = "./aws_iam_user_login_profile"
  user                    = var.user
  pgp_key                 = var.pgp_key
  password_length         = var.password_length
  password_reset_required = var.password_reset_required
}

module "aws_ssm_parameter" {
  source = "./aws_ssm_parameter"
  name   = var.name
  value  = var.value
}

module "aws_secretsmanager_secret" {
  source = "./aws_secretsmanager_secret"
  name   = var.name
}

module "aws_secretsmanager_secret_version" {
  source        = "./aws_secretsmanager_secret_version"
  secret_id     = var.secret_id
  secret_string = var.secret_string
  secret_binary = var.secret_binary
}

### Bucket

module "aws_s3_bucket" {
  source = "./aws_s3_bucket"
  bucket = var.bucket
}
