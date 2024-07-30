variable "secret_id" {
  description = "Specifies the secret to which you want to add a new version"
  type        = string
  default     = null
}

variable "secret_string" {
  description = "Specifies text data that you want to encrypt and store in this version of the secret"
  type        = string
  default     = null
}

variable "secret_binary" {
  description = "Specifies binary data that you want to encrypt and store in this version of the secret"
  type        = string
  default     = null
}
