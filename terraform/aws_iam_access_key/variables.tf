variable "user" {
  description = "IAM user to associate with this access key"
  type        = string
}

variable "pgp_key" {
  description = "For use in the encrypted_secret output attribute"
  type        = string
  default     = null
}

variable "status" {
  description = "Access key status to apply"
  type        = string
  default     = null
}
