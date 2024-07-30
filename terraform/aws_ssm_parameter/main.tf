resource "aws_ssm_parameter" "aws_ssm_parameter" {
  name  = var.name
  type  = var.type
  value = var.value
}
