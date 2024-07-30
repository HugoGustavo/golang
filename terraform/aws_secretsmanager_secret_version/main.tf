resource "aws_secretsmanager_secret_version" "aws_secretsmanager_secret_version" {
    secret_id     = var.secret_id
    secret_string = var.secret_string
}
