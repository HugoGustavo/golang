variable "bucket" {
  description = "Name of the bucket"
  type        = string
}

variable "force_destroy" {
  description = "All objects should be deleted from the bucket when the bucket is destroyed"
  type        = string
  default     = true
}
