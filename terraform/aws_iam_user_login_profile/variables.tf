variable "user" {
  description = "The IAM user's name"
  type        = string
}

variable "pgp_key" {
  description = "Either a base-64 encoded PGP public key"
  type        = string
  default     = null
}

variable "password_length" {
  description = "The length of the generated password on resource creation"
  type        = number
  default     = 20
}

variable "password_reset_required" {
  description = "Whether the user should be forced to reset the generated password on resource creation"
  type        = bool
  default     = true
}
