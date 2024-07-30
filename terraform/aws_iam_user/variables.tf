variable "name" {
  description = "The user's name"
  type        = string
}

variable "path" {
  description = "Path in which to create the user"
  type        = string
  default     = "/"
}

variable "permissions_boundary" {
  description = "The ARN of the policy that is used to set the permissions boundary for the user"
  type        = string
  default     = null
}

variable "force_destroy" {
  description = "Destroy even if it has non-Terraform-managed IAM access keys, login profile or MFA devices"
  type        = bool
  default     = true
}

variable "tags" {
  description = "Key-value map of tags for the IAM user"
  type        = map(any)
  default     = null
}
