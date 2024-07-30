variable "name" {
  description = "The group's name"
  type        = string
}

variable "path" {
  description = "Path in which to create the group"
  type        = string
  default     = "/"
}